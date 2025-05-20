package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "KubeGuard",
	Short: "Kubernetes 应用与节点数据备份恢复工具",
	Long:  "KubeGuard 是一个用于备份和恢复 Kubernetes 应用资源与节点数据的 CLI 工具。",
}

func Execute() error {
	rootCmd.AddCommand(backupCmd, restoreCmd, listCmd, deleteCmd)
	return rootCmd.Execute()
}

