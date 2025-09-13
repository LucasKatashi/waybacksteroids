package core

import "fmt"

func PrintOutput(urls []string) {
	for _, url := range urls {
		fmt.Println(url)
	}
}
