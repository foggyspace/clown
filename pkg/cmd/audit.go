package cmd

import (
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "select vulnerabilty type ?",
	Long:  "select vulnerabilty type from list ?",
}

func init() {
	rootCmd.AddCommand(auditCmd)
}
