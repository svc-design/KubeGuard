# Makefile for KubeGuard

BINARY_NAME=k8s_backup_tool
BUILD_DIR=bin
CONFIG_FILE=k8s_backup_config.yaml

VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-X main.version=$(VERSION) -s -w"

.PHONY: all build run clean test fmt vet release linux-amd64 linux-arm64

all: build

init:
	GOPROXY=https://goproxy.cn,direct go get github.com/spf13/cobra@latest
	go mod tidy

# 构建本地二进制
build:
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go

# 运行工具 (默认展示帮助信息)
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) --help

# 格式化代码
fmt:
	go fmt ./...

# 检测代码潜在错误
vet:
	go vet ./...

# 执行单元测试
test:
	go test ./... -cover

# 清理构建产物
clean:
	rm -rf $(BUILD_DIR)/*

# 交叉编译 Linux amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 cmd/main.go

# 交叉编译 Linux arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 cmd/main.go

# 同时构建 amd64 和 arm64 二进制包
release: clean linux-amd64 linux-arm64
	tar -czvf $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-amd64
	tar -czvf $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-arm64

