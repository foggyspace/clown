package cmd

import (
	"github.com/spf13/cobra"
)

var auditor = &cobra.Command{
	Use:   "audit",
	Short: "请选择审计的类型",
	Long:  "这个命令提供了一些子命令，您可根据这些子命令选择需要的漏洞进行审计",
}

func init() {
	rootCmd.AddCommand(auditor)
}
