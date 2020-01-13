package engine

import (
	"fmt"
	"log"
	"Cyker/utils"
)

func Run(targetURL string) {
	log.Fatal(targetURL)
	utils.Banner()
	fmt.Println("[+] target URL : ", targetURL)
}
