package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/seaung/Cyker/commons"
	"github.com/seaung/Cyker/lib"
)

func initProgram() {
	commons.CheckSudo()

	commons.Banner()
}

func main() {
	initProgram()

	p := prompt.New(
		lib.Executor,
		lib.Completer,
		prompt.OptionTitle("Cyker: A Poc tools"),
		prompt.OptionPrefix("[Cyker] > "),
		prompt.OptionInputTextColor(prompt.White),
	)
	p.Run()
}
