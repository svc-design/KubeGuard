func BackupNodeData(nodeName, srcPath, dateTag string, cfg *config.Config) error {
  snapshotDir := filepath.Join("/tmp/k8s-node-backup", nodeName)
  os.MkdirAll(snapshotDir, 0755)

  utils.ExecCommand("rsync", "-aHAX", "--numeric-ids", srcPath+"/", snapshotDir)

  archivePath := fmt.Sprintf("/tmp/%s_backup.tar.xz", nodeName)
  utils.ExecCommand("tar", "--preserve-permissions", "--same-owner", "-cJf", archivePath, "-C", "/tmp/k8s-node-backup", nodeName)

  sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.Settings.VeleroRegion)}))
  key := fmt.Sprintf("%s/%s/%s_backup.tar.xz", cfg.BackupConfig.ClusterName, dateTag, nodeName)

  storage.UploadFile(sess, cfg.Settings.VeleroBucket, key, archivePath)
  return nil
}

