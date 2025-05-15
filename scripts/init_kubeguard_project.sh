#!/bin/bash
# init_kubeguard_project.sh
# 初始化 KubeGuard 项目目录结构

set -e

# 创建项目目录结构
mkdir -p cmd \
    pkg/{config,backup,restore,storage,k8s,node,utils} \
    scripts \
    .github/workflows

# 创建基础文件
touch cmd/main.go
touch pkg/config/config.go
touch pkg/backup/backup.go
touch pkg/restore/restore.go
touch pkg/storage/s3.go
touch pkg/k8s/velero.go
touch pkg/node/node.go
touch pkg/utils/{command.go,logger.go}
touch scripts/offline_install.sh
touch .github/workflows/release.yml
touch README.md

# 初始化 go.mod
go mod init github.com/svc-design/KubeGuard

echo "✅ KubeGuard 项目结构初始化完成！"
