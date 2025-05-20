package cmd

import (
	"kubeguard/internal/shell"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "创建 K8s 应用资源备份 ➕ 节点数据打包并上传 S3",
	Run: func(cmd *cobra.Command, args []string) {
		shell.RunCmd("backup")
	},
}

