package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/LucasKatashi/waybacksteroids/internal/core"
	"github.com/LucasKatashi/waybacksteroids/internal/ui"
	"github.com/fatih/color"
)

var (
	red = color.New(color.FgRed)
)

func main() {
	target := flag.String("target", "", "")
	flag.StringVar(target, "t", "", "")

	threads := flag.Int("threads", 3, "")

	retries := flag.Int("retries", 3, "")
	flag.IntVar(retries, "r", 3, "")

	wordlist := flag.String("wordlist", "", "")
	flag.StringVar(wordlist, "w", "", "")

	output := flag.String("output", "", "")
	flag.StringVar(output, "o", "", "")

	silent := flag.Bool("silent", false, "")
	flag.BoolVar(silent, "s", false, "")

	verbose := flag.Bool("verbose", false, "")
	flag.BoolVar(verbose, "v", false, "")

	print := flag.Bool("print", false, "")
	flag.BoolVar(print, "p", false, "")

	stdin := flag.Bool("stdin", false, "")

	flag.Usage = ui.CustomUsage
	flag.Parse()

	ui.PrintBanner(*silent)

	var result []string
	var domainCount int

	if *threads > 5 {
		fmt.Printf("[%s] Using more than 5 threads may get your IP address blocked.\n", color.YellowString("WRN"))
	}

	if *stdin {
		scanner := bufio.NewScanner(os.Stdin)
		var domains []string
		for scanner.Scan() {
			domain := scanner.Text()
			if domain != "" {
				domains = append(domains, domain)
			}
		}
		if err := scanner.Err(); err != nil {
			return
		}
		domainCount = len(domains)
		result = core.DomainsParser(domains)
	} else if *target != "" {
		domainCount = 1
		result = core.DomainsParser([]string{*target})
	} else if *wordlist != "" {
		file := core.ReadWordlist(*wordlist)
		domainCount = len(file)
		result = core.DomainsParser(file)
	} else {
		fmt.Printf("[%s] You need to specify either --target/-t, --wordlist/-w, or --stdin\n", red.Sprint("ERR"))
		return
	}

	fmt.Printf("[%s] Loaded %d domain(s)\n", color.BlueString("INF"), domainCount)

	if *output == "" {
		fmt.Printf("[%s] Output directory is required. Use --output/-o to specify a directory\n", red.Sprint("ERR"))
		return
	}

	core.WebArchiveRequest(result, *verbose, *print, *threads, *output, *retries)
}
