package commons

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

var (
	author string
	version string
)


func Banner() {
	name := fmt.Sprintf("Cyker (v.%s)", version)
	logo := `
				  _               
	  ___  _   _ | | __ ___  _ __ 
	 / __|| | | || |/ // _ \| '__|
	| (__ | |_| ||   <|  __/| |   
	 \___| \__, ||_|\_\\___||_|   
		   |___/   
	`

	lines := strings.Split(logo, "\n")
	w := len(lines[1])

	fmt.Println(logo)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w + len(name)) / 2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w + len(author)) / 2, author)))
	fmt.Println()
}

