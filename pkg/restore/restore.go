func Run(cfg *config.Config, dateTag string) error {
  utils.ExecShell(cfg.BackupConfig.PreCmds)

  for node, destPath := range cfg.BackupConfig.Nodes {
    node.RestoreNodeData(node, destPath, dateTag, cfg)
  }

  veleroBackupName, _ := k8s.GetBackupNameByLabel(cfg, dateTag)
  k8s.RestoreFromBackup(veleroBackupName, cfg.Settings.VeleroNamespace)

  utils.ExecShell(cfg.BackupConfig.PostCmds)
  return nil
}
