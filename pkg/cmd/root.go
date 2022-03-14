package cmd

import (
	"fmt"
	"os"

	"github.com/seaung/clown/pkg/utils"
	"github.com/spf13/cobra"
)

var longTxt = `
        __                  
  _____/ /___ _      ______ 
 / ___/ / __ \ | /| / / __ \
/ /__/ / /_/ / |/ |/ / / / /
\___/_/\____/|__/|__/_/ /_/ 

        Version: 1.1.1
        小丑竟然是我自己?
	Author: [ seaung ]
	Github: https://github.com/seaung
`

var rootCmd = &cobra.Command{
	Use:   "clown",
	Short: "一个交互式命令行工具",
	Long:  longTxt,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.NewLogger().Warnning(fmt.Sprintf("A valid command %v\n", err))
		os.Exit(1)
	}
}
