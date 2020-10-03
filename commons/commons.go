package commons

import (
	"os"
	"strings"
)

var Const_example_target_cidr = "127.0.0.1/32"
var Const_example_target_desc = "Target CIDR or /32 for single target"

func ParseCmd(s string) (string, []string) {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "", make([]string, 0)
	}

	tokens := strings.Fields(s)

	cmd, args := tokens[0], tokens[1:]
	return cmd, args
}

func ParseNextArg(args []string) (string, []string) {
	if len(args) < 2 {
		return "", make([]string, 0)
	}
	return args[0], args[1:]
}

func ParseAllArgs(args []string) string {
	return strings.Join(args, " ")
}

func CheckSudo() {
	if os.Getuid() != 0 {
		os.Exit(1)
	}
}

