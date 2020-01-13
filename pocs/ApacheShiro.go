package pocs

import (
  "fmt"
)

type ApacheShiro struct {}

func (a *ApacheShiro) Audit(targetURL string) {
  fmt.Println(a)
}
