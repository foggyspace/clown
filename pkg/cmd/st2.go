package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

/*
audit下的子命令,专门为st2漏洞设定的一个漏洞检测命令
这个子命令提供了一些常见的st2漏洞列表供您选择 ~_~
使用的命令格式，如下所示:
root > clown audit st2 s2-057 https://www.vulntest.com.cn
*/
var st2Cmd = &cobra.Command{
	Use:   "st2",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	auditCmd.AddCommand(st2Cmd)
}

func auditSt2(target string) {
}

func selectPromptForSt2(pc promptContent) string {
	var err error
	var selectType string
	items := []string{"s2-001", "s2-005", "s2-007", "s2-008", "s2-009", "s2-012", "s2-013", "s2-052", "s2-057"}
	prompt := promptui.SelectWithAdd{
		Label:    pc.label,
		Items:    items,
		AddLabel: "Other",
	}

	_, selectType, err = prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("你选择检测的检测的漏洞是 : %s\n", selectType)
	return selectType
}
