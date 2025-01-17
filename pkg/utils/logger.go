package utils

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Warnning(message string) {
	highlight := color.New(color.FgYellow).SprintFunc()
	fmt.Println(highlight("[!]"), highlight(message))
}

func (l *Logger) ErrorLog(message string) {
	highlight := color.New(color.FgRed).SprintFunc()
	fmt.Println(highlight("[!]"), highlight(message))
}

func (l *Logger) Info(message string) {
	highlight := color.New(color.FgBlue).SprintFunc()
	reset := color.New(color.FgWhite).SprintFunc()
	fmt.Println(highlight("[*]"), reset(message))
}
