package lib

import (
	"fmt"
	"os"
	"github.com/seaung/Cyker/commons"
)

func Executor(s string) {
	cmd, _ := commons.ParseCmd(s)

	switch cmd {
	case "verify":
		fmt.Println("verify")
	case "help":
		commons.ShowHelp()
	case "exit", "quit":
		os.Exit(0)
		return
	case "":
	default:
		return
	}
}
