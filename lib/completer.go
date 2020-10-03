package lib

import (
	"strings"
	"github.com/c-bata/go-prompt"
	"github.com/seaung/Cyker/commons"
)

var commands = []prompt.Suggest{
	{Text: "verifiy", Description: "verifiy vulnerable"},
	{Text: "help", Description: "show help"},
	{Text: "exit", Description: "Exit this program"},
	{Text: "show", Description: "show results"},
}

func excludeOptions(args []string) []string {
	ret := make([]string, 0, len(args))
	for i := range args {
		if !strings.HasPrefix(args[i], "-") {
			ret = append(ret, args[i])
		}
	}
	return ret
}

func argumentsCompleter(d prompt.Document, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]

	switch first {
	case "verifiy":
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "ALL", Description: "Use all POCs for detection."},
				{Text: "JAVA", Description: "Java vulnerability detection only."},
				{Text: "PHP", Description: "PHP vulnerability detection only"},
			}
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
		if len(args) == 3 {
			return prompt.FilterContains(commons.GetTargetSuggestions(), args[2], true)
		}
	}
	return []prompt.Suggest{}
}

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := strings.Split(d.TextBeforeCursor(), " ")

	for i:= range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}
	return argumentsCompleter(d, excludeOptions(args))
}
