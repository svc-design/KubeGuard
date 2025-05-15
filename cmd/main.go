package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/svc-design/KubeGuard/pkg/backup"
	"github.com/svc-design/KubeGuard/pkg/config"
	"github.com/svc-design/KubeGuard/pkg/restore"
)

var cfgFile string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "k8s_backup_tool",
		Short: "📘 Kubernetes 集群备份与恢复工具 (v4.15.16)",
		Long: `📘 Kubernetes 集群备份与恢复工具 (v4.15.16)

命令        说明
backup      创建 K8s 应用资源备份 ➕ 节点数据打包并上传 S3
restore     先恢复节点数据，再恢复 Velero 应用资源
list        列出所有备份（Velero + S3），自动对齐 date_tag
delete      删除指定 date_tag 的 Velero + S3 备份

示例：
  k8s_backup_tool list -c k8s_backup_config.yaml
  k8s_backup_tool backup -c k8s_backup_config.yaml
  k8s_backup_tool delete -c k8s_backup_config.yaml <date_tag>
  k8s_backup_tool restore -c k8s_backup_config.yaml <date_tag>`}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (required)")
	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "创建 K8s 应用资源备份 ➕ 节点数据打包并上传 S3",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("❌ 配置加载失败: %v", err)
		}
		err = backup.Run(cfg)
		if err != nil {
			log.Fatalf("❌ 备份失败: %v", err)
		}
	},
}

var restoreCmd = &cobra.Command{
	Use:   "restore <date_tag>",
	Short: "先恢复节点数据，再恢复 Velero 应用资源",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("❌ 配置加载失败: %v", err)
		}
		err = restore.Run(cfg, args[0])
		if err != nil {
			log.Fatalf("❌ 恢复失败: %v", err)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有备份（Velero + S3），自动对齐 date_tag",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("❌ 配置加载失败: %v", err)
		}
		err = backup.List(cfg)
		if err != nil {
			log.Fatalf("❌ 列出备份失败: %v", err)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <date_tag>",
	Short: "删除指定 date_tag 的 Velero + S3 备份",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			log.Fatalf("❌ 配置加载失败: %v", err)
		}
		err = backup.Delete(cfg, args[0])
		if err != nil {
			log.Fatalf("❌ 删除备份失败: %v", err)
		}
	},
}
