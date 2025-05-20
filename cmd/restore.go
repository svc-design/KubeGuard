package cmd

import (
	"kubeguard/internal/shell"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore <tag>",
	Short: "恢复节点数据 ➕ 应用资源",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell.RunCmd("restore", args[0])
	},
}

