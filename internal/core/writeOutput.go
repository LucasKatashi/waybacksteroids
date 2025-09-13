package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteOutput(outputDir string, domain string, urls []string) {
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("[%s] Error creating directory:\n      %s\n", red.Sprint("ERR"), err)
		return
	}

	sanitizedDomain := strings.ReplaceAll(domain, "/", "_")

	filename := sanitizedDomain + "_steroids.txt"
	filepath := filepath.Join(outputDir, filename)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("[%s] Error creating file:\n      %s\n", red.Sprint("ERR"), err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, url := range urls {
		writer.WriteString(url + "\n")
	}
	writer.Flush()
}
