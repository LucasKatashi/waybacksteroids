package core

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/LucasKatashi/waybacksteroids/internal/ui"
)

var (
	connectionRefusedCount int
	mutex                  sync.Mutex
	blockMessageShown      bool
)

func WebArchiveRequest(domains []string, verbose bool, print bool, threads int, output string, retries int) {
	sem := make(chan struct{}, threads)
	var wg sync.WaitGroup

	for _, domain := range domains {
		wg.Add(1)
		go func(d string) {
			sem <- struct{}{}
			defer func() {
				<-sem
			}()
			defer wg.Done()

			url := fmt.Sprintf("http://web.archive.org/cdx/search/cdx?url=*.%s*&output=text&fl=original", d)

			var resp *http.Response
			var err error
			var body []byte

			for attempt := 1; attempt <= retries; attempt++ {
				resp, err = http.Get(url)
				if err != nil {
					if attempt == retries {
						if strings.Contains(err.Error(), "connection refused") {
							mutex.Lock()
							connectionRefusedCount++
							if connectionRefusedCount >= 10 && !blockMessageShown {
								blockMessageShown = true
								mutex.Unlock()
								ui.PrintIPblockMessage(threads)
							}
							mutex.Unlock()
						}
						fmt.Printf("[%s] Failed to connect to Wayback Machine after %d attempts:\n      %s\n", color.RedString("ERR"), retries, err)
						return
					}

					time.Sleep(5 * time.Second)
					continue
				}

				if resp.StatusCode == 200 {
					break
				}

				resp.Body.Close()

				if attempt < 3 {
					time.Sleep(5 * time.Second)
				}

				if attempt == retries {
					fmt.Printf("[%s] Wayback Machine returned status %d for %s after %d attempts\n", color.RedString("ERR"), resp.StatusCode, d, retries)
					return
				}
			}

			if err != nil || resp.StatusCode != 200 {
				return
			}

			defer resp.Body.Close()

			body, err = io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("[%s] Failed to read response body:\n      %s\n", color.RedString("ERR"), err)
				return
			}

			lines := strings.Split(string(body), "\n")
			seen := make(map[string]bool)
			var urls []string

			for _, line := range lines {
				if line != "" && !seen[line] {
					urls = append(urls, line)
					seen[line] = true
				}
			}

			if print {
				PrintOutput(urls)
			}

			if output != "" {
				WriteOutput(output, d, urls)
				if !print {
					fmt.Printf("[%s] Saved %d archived URLs for:\n     %s\n", color.GreenString("OK"), len(urls), d)
				}
			}
		}(domain)
	}

	wg.Wait()

	fmt.Printf("[%s] Processing completed!\n", color.GreenString("OK"))
}
