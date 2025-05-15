func ExecCommand(name string, args ...string) (string, error) {
  cmd := exec.Command(name, args...)
  out, err := cmd.CombinedOutput()
  return string(out), err
}

func ExecShell(cmdStr string) error {
  cmd := exec.Command("bash", "-c", cmdStr)
  return cmd.Run()
}
