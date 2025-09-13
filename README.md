<h4 align="center"><b>waybacksteroids</b> — Fast multi-domain Wayback Machine endpoint enumerator.</h4>

<p align="center">
  <a href="#installation-instructions">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#running-waybacksteroids">Running waybacksteroids</a>
</p>

---

**waybacksteroids** is a enumeration tool that automates the retrieval of archived URLs from the Wayback Machine. It supports processing multiple domains simultaneously, making it useful for quickly discovering historical endpoints and uncovering hidden paths across different targets.

# Installation Instructions
```sh
go install github.com/LucasKatashi/waybacksteroids/cmd/waybacksteroids@latest
```

# Usage
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

EXAMPLES:
 # Single domain
 waybacksteroids -t example.com -o output/

 # Multiple domains from file
 waybacksteroids -w domains.txt -o output/

 # Pipe from subdomain enumeration tool
 subfinder -d example.com | waybacksteroids --stdin -o output/

 # Print to stdout only
 waybacksteroids -t example.com -p
```

---

**waybacksteroids** 是一款自动化从Wayback Machine获取存档URL的枚举工具。它支持同时处理多个域名，能够快速发现历史端点并揭示不同目标中隐藏的路径，非常实用。

# 安装说明
```sh
go install github.com/LucasKatashi/waybacksteroids/cmd/waybacksteroids@latest
```

# 使用方法
```sh
waybacksteroids -h
```

这将显示工具的帮助信息。以下是所有支持的选项。
```console
用法:
 waybacksteroids [选项]

输入方式 (选择一种):
 -t, --target           单个目标域名 (例如：example.com)
 -w, --wordlist         包含域名列表的文件 (每行一个域名)
 --stdin                从标准输入读取域名 (从其他工具管道传输)

配置选项:
 -o, --output           输出目录 (必需) - 结果将保存为 domain_steroids.txt 文件
 --threads              并发线程数 (默认：3，推荐最大值：5)
 -r, --retries          失败请求的重试次数 (默认：3)

输出选项:
 -p, --print            将URL打印到标准输出 (不创建文件)
 -v, --verbose          启用详细模式
 -s, --silent           隐藏横幅和信息消息

示例:
 # 单个域名
 waybacksteroids -t example.com -o output/

 # 从文件读取多个域名
 waybacksteroids -w domains.txt -o output/

 # 从子域名枚举工具管道传输
 subfinder -d example.com | waybacksteroids --stdin -o output/

 # 仅打印到标准输出
 waybacksteroids -t example.com -p
```
