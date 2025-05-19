package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func RunScript(operation string, args ...string) {
	script := "scripts/k8s_backup_tool.sh"
	allArgs := append([]string{script, operation, "-c", "k8s_backup_config.yaml"}, args...)

	cmd := exec.Command("bash", allArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🔧 执行命令: bash %v\n", allArgs)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ 脚本执行失败: %v\n", err)
		os.Exit(1)
	}
}

