// internal/shell/exec.go
package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCmd executes a shell command with arguments and streams its output.
// Returns an error if execution fails.
func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fullCmd := fmt.Sprintf("%s %s", name, strings.Join(args, " "))
	fmt.Printf("🔧 执行命令: %s\n", fullCmd)

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ 命令执行失败: %v\n", err)
		return err
	}
	return nil
}

