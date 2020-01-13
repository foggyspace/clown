package main

import (
  "os"
  "Cyker/engine"
  "Cyker/utils"
)

func main() {
  if len(os.Args) != 2 {
    utils.Banner()
  }
  engine.Run(os.Args[1])
}
