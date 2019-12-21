package main

import (
	"fmt"
	"os"

	"Cyker/lib"
	"Cyker/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[+] Usage : ", os.Args[0], "<url>")
		os.Exit(0)
	}
	targetURL := os.Args[1]
	utils.Banner()
	crawler := lib.WCrawler{}
	crawler.Crawler(targetURL)

	for _, link := range crawler.AllLinks {
		fmt.Println("[*] Found : ", link)
	}
}
