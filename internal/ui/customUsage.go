package ui

import (
	"fmt"
	"os"
)

func CustomUsage() {
	fmt.Fprintf(os.Stderr, "waybacksteroids - Wayback Machine enumeration tool\n")
	fmt.Fprintf(os.Stderr, "Retrieves archived URLs from the Wayback Machine for specified domains.\n\n")

	fmt.Fprintf(os.Stderr, "Usage:\n waybacksteroids [flags]\n\n")

	fmt.Fprintf(os.Stderr, "INPUT METHODS (choose one):\n")
	fmt.Fprintf(os.Stderr, " -t, --target\t\tsingle target domain (e.g., example.com)\n")
	fmt.Fprintf(os.Stderr, " -w, --wordlist\t\tfile containing list of domains (one per line)\n")
	fmt.Fprintf(os.Stderr, " --stdin\t\tread domains from stdin (pipe from other tools)\n\n")

	fmt.Fprintf(os.Stderr, "CONFIGURATION:\n")
	fmt.Fprintf(os.Stderr, " -o, --output\t\toutput directory (required) - saves results in domain_steroids.txt files\n")
	fmt.Fprintf(os.Stderr, " --threads\t\tnumber of concurrent threads (default: 3, max recommended: 5)\n")
	fmt.Fprintf(os.Stderr, " -r, --retries\t\tnumber of retry attempts for failed requests (default: 3)\n\n")

	fmt.Fprintf(os.Stderr, "OUTPUT OPTIONS:\n")
	fmt.Fprintf(os.Stderr, " -p, --print\t\tprint URLs to stdout (no files created)\n")
	fmt.Fprintf(os.Stderr, " -v, --verbose\t\tenable verbose mode\n")
	fmt.Fprintf(os.Stderr, " -s, --silent\t\tsuppress banner and info messages\n\n")

	fmt.Fprintf(os.Stderr, "EXAMPLES:\n")
	fmt.Fprintf(os.Stderr, " # Single domain\n waybacksteroids -t example.com -o output/\n\n")
	fmt.Fprintf(os.Stderr, " # Multiple domains from file\n waybacksteroids -w domains.txt -o output/\n\n")
	fmt.Fprintf(os.Stderr, " # Pipe from subdomain enumeration tool\n seekly -domain example.com | waybacksteroids --stdin -o output/\n\n")
	fmt.Fprintf(os.Stderr, " # Print to stdout only\n waybacksteroids -t example.com -p\n")
}
