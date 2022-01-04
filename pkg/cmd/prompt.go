package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

func getPromptInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	res, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return res
}

func getPromptSelect(pc promptContent) string {
	prompt := promptui.Select{
		Label: pc.label,
		Items: []string{"java", "common"},
	}
	_, res, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}
