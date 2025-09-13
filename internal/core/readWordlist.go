package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	red = color.New(color.FgRed)
)

func ReadWordlist(wordlist string) []string {
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("[%s] Error opening wordlist file:\n      %s\n", red.Sprint("ERR"), err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[%s] Error reading wordlist file:\n      %s\n", red.Sprint("ERR"), err)
		os.Exit(1)
	}

	return lines
}
