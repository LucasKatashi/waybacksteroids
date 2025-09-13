<h4 align="center"><b>waybacksteroids</b> — Fast multi-domain Wayback Machine endpoint enumerator.</h4>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation-instructions">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#running-waybacksteroids">Running waybacksteroids</a> •
  <a href="#examples">Examples</a>
</p>

---

**waybacksteroids** is an enumeration tool that automates the retrieval of archived URLs from the Wayback Machine. It supports processing **multiple domains simultaneously**, making it useful for quickly discovering historical endpoints and uncovering hidden paths across different targets.

## Features
- 🔍 Bulk fetch endpoints for any number of domains in one run  
- 🚀 Concurrent requests (configurable threads, default 3, **max 5**)  
- 📤 Flexible input: single domain, wordlist or stdin pipe (fits any recon workflow)  
- 🎯 Clean, de-duplicated output per domain (`domain_steroids.txt`)  
- 🖥️  Print-only mode for quick terminal checks (`-p`)  
- 🔁  Auto-retry on transient failures (configurable)  
- 📦  Single static binary—no dependencies after compile

## Installation Instructions

### From source (Go ≥ 1.20)
```sh
go install github.com/LucasKatashi/waybacksteroids/cmd/waybacksteroids@latest
```

### Pre-built binaries

### Build manually
```sh
git clone https://github.com/LucasKatashi/waybacksteroids.git
cd waybacksteroids
go build -o waybacksteroids ./cmd/waybacksteroids
```

## Usage
```sh
waybacksteroids -h
```

This will display help for the tool. Here are all the switches it supports.
```console
Usage:
 waybacksteroids [flags]

INPUT METHODS (choose one):
 -t, --target           single target domain (e.g., example.com)
 -w, --wordlist         file containing list of domains (one per line)
 --stdin                read domains from stdin (pipe from other tools)

CONFIGURATION:
 -o, --output           output directory (required) - saves results in domain_steroids.txt files
 --threads              number of concurrent threads (default: 3, max recommended: 5)
 -r, --retries          number of retry attempts for failed requests (default: 3)

OUTPUT OPTIONS:
 -p, --print            print URLs to stdout (no files created)
 -v, --verbose          enable verbose mode
 -s, --silent           suppress banner and info messages
```

## Examples

**Single domain**  
```sh
waybacksteroids -t example.com -o output/
```

**Multiple domains from file**  
```sh
waybacksteroids -w domains.txt -o output/
```

**Pipe from subdomain enumeration tool**  
```sh
subfinder -d example.com | waybacksteroids --stdin -o output/
```

**Print to stdout only**  
```sh
waybacksteroids -t example.com -p
```

## Sample Output
Running  
```sh
waybacksteroids -t example.com -o output/
```
creates `output/example.com_steroids.txt` containing:
```
http://example.com/robots.txt
http://example.com/.git/config
http://example.com/api/v1/users
http://example.com/admin/login.jsp
[...]
```
