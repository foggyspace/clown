package utils

import (
	"fmt"
)

const (
	AUTHOR  = "seaung"
	VERSION = "1.0.0"
)

func Banner() {
	logo := `
	+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+
	|c|y|k|e|r| |w|e|b|s|a|n|n|e|r|
	+-+-+-+-+-+ +-+-+-+-+-+-+-+-+-+
	`
	fmt.Println(logo)
	fmt.Println(AUTHOR)
	fmt.Println(VERSION)
}
