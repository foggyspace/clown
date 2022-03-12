package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/seaung/clown/pkg/plugins"
	"github.com/spf13/cobra"
)

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Java 类型的通用型的漏洞",
	Long:  "请选择一个通用型java漏洞插件",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	auditor.AddCommand(javaCmd)
}

func run() {
	targetPrompt := clownPromptContent{
		"请您提供一个需要检测的目标URL!",
		"请您输入一个目标: ",
	}

	target := getPromptInput(targetPrompt)

	selectGateroiesPrompt := clownPromptContent{
		"请您提供一个需要检测的Java类型漏洞",
		"请您选择: ",
	}

	gategories := getSelectJavaCategoriesPrompt(selectGateroiesPrompt)

	fmt.Println("[*] 需要检测的目标为  : ", target)
	fmt.Println("[*] 您选择的漏洞类型为: ", gategories)
	runVulns(target, gategories)
}

func runVulns(target, gategories string) {
	switch gategories {
	default:
		fmt.Println("您是不是忘记选择漏洞类型了?")
	case "jboss":
		jboss := &plugins.JbossCve201712149{}
		jboss.Audit(target)
	}
}

func getSelectJavaCategoriesPrompt(cpc clownPromptContent) string {
	items := []string{"jboss", "active"}
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
