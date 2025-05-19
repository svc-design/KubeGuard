# KubeGuard

Kubernetes 应用与节点数据备份恢复 CLI 工具。
KubeGuard: 🚀 Easy Kubernetes cluster backup and restore tool, built with Velero and optimized for containerized environments.

# KubeGuard 目录结构 
```
├── .gitignore
├── Makefile
├── Readme.md
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   ├── root.go
│   ├── backup.go
│   ├── restore.go
│   ├── list.go
│   └── delete.go
├── internal/
│   └── shell/
│       └── exec.go
└── scripts/
    └── k8s_backup_tool.sh
```

# TLDR

Usage:
  KubeGuard [command] -c <ConfigPath>

Available Commands:
  backup      创建 K8s 应用资源备份 ➕ 节点数据打包并上传 S3
  completion  Generate the autocompletion script for the specified shell
  delete      删除指定备份
  help        Help about any command
  list        列出所有备份（Velero + S3）
  restore     恢复节点数据 ➕ 应用资源

Flags:
  -h, --help   help for k8s_backup_tool
