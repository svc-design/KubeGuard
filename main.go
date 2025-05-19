package main

import (
	"fmt"
	"os"

	"kubeguard/cmd"
)

func main() {
	fmt.Println("📘 使用说明：k8s_backup_tool v4.15.16")
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ 错误: %v\n", err)
		os.Exit(1)
	}
}

