package cmd

import (
	"kubeguard/internal/shell"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有备份（Velero + S3）",
	Run: func(cmd *cobra.Command, args []string) {
		shell.RunScript("list")
	},
}
