package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type clownPromptContent struct {
	errorMsg string
	label    string
}

func getPromptInput(cpc clownPromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(cpc.errorMsg)
		}
		return nil
	}

	promptTemplates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | blue }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     cpc.label,
		Templates: promptTemplates,
		Validate:  validate,
	}

	content, err := prompt.Run()
	if err != nil {
		fmt.Printf("[!] Prompt Error %v\n", err)
		os.Exit(1)
	}

	return content
}

func getSelectPrompt(cpc clownPromptContent) string {
	items := []string{"java", "common"}
	index := -1

	var content string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    cpc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, content, err = prompt.Run()

		if index == -1 {
			items = append(items, content)
		}
	}

	if err != nil {
		fmt.Printf("Prompt Select Error %v\n", err)
		os.Exit(1)
	}

	return content
}
