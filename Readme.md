# KubeGuard

Kubernetes 应用与节点数据备份恢复 CLI 工具。
KubeGuard: 🚀 Easy Kubernetes cluster backup and restore tool, built with Velero and optimized for containerized environments.

KubeGuard/
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
