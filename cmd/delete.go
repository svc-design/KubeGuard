package cmd

import (
	"kubeguard/internal/shell"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <tag>",
	Short: "删除指定备份",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell.RunCmd("delete", args[0])
	},
}

