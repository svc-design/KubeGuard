APP_NAME := KubeGuard
MAIN_FILE := main.go
TAG ?= latest
CONFIG ?= k8s_backup_config.yaml

.PHONY: all build run clean init backup restore delete list help

all: build

init:
	GOPROXY=https://goproxy.cn,direct go get github.com/spf13/cobra@latest
	go mod tidy

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE) list

backup:
	go run $(MAIN_FILE) backup

restore:
	go run $(MAIN_FILE) restore $(TAG)

delete:
	go run $(MAIN_FILE) delete $(TAG)

list:
	go run $(MAIN_FILE) list

clean:
	rm -f $(APP_NAME)

help:
	@echo "🔧 KubeGuard CLI Usage"
	@echo ""
	@echo "make build           编译可执行文件"
	@echo "make run             默认执行 list"
	@echo "make list            列出所有备份"
	@echo "make backup          创建备份"
	@echo "make restore TAG=xxx 恢复指定备份"
	@echo "make delete TAG=xxx  删除指定备份"
	@echo "make init            初始化依赖 (go mod tidy)"
	@echo "make clean           清理构建产物"
