package core

import "strings"

func DomainsParser(domains []string) []string {
	processed := make([]string, 0)
	seen := make(map[string]bool)

	for _, domain := range domains {
		domain = strings.TrimPrefix(domain, "https://")
		domain = strings.TrimPrefix(domain, "http://")
		domain = strings.TrimSuffix(domain, "/")
		domain = domain + "/"

		if !seen[domain] {
			seen[domain] = true
			processed = append(processed, domain)
		}
	}

	return processed
}
