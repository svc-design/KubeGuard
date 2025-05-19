package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s_backup_tool",
	Short: "Kubernetes 应用与节点数据备份恢复工具",
	Long: `📘 使用说明：k8s_backup_tool v4.15.16

命令        说明
backup         创建 K8s 应用资源备份 ➕ 节点数据打包并上传 S3
restore <tag>  先恢复节点数据，再恢复 Velero 应用资源
list           列出所有备份（Velero + S3），自动对齐 date_tag
delete <tag>   删除指定 date_tag 的 Velero + S3 备份

示例：
  bash scripts/k8s_backup_tool.sh list -c k8s_backup_config.yaml
  bash scripts/k8s_backup_tool.sh backup -c k8s_backup_config.yaml
  bash scripts/k8s_backup_tool.sh delete -c k8s_backup_config.yaml <date_tag>
  bash scripts/k8s_backup_tool.sh restore -c k8s_backup_config.yaml <date_tag>`,
}

func Execute() error {
	rootCmd.AddCommand(backupCmd, restoreCmd, listCmd, deleteCmd)
	return rootCmd.Execute()
}
