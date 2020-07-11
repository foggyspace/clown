package commons

import (
	"fmt"
)

var logo string = `
              _               
  ___  _   _ | | __ ___  _ __ 
 / __|| | | || |/ // _ \| '__|
| (__ | |_| ||   <|  __/| |   
 \___| \__, ||_|\_\\___||_|   
       |___/   
`

func Banner() {
	fmt.Println(logo)
}
