type Config struct {
  Settings struct {
    VeleroNamespace   string `mapstructure:"VELERO_NAMESPACE"`
    VeleroBucket      string `mapstructure:"VELERO_BUCKET"`
    VeleroRegion      string `mapstructure:"VELERO_REGION"`
    AWSAccessKeyID    string `mapstructure:"AWS_ACCESS_KEY_ID"`
    AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
  }
  BackupConfig struct {
    ClusterName string
    Nodes       map[string]string
    Namespaces  []string
    PreCmds     string
    PostCmds    string
  }
}

func LoadConfig(path string) (*Config, error) {
  v := viper.New()
  v.SetConfigFile(path)

  if err := v.ReadInConfig(); err != nil {
    return nil, err
  }

  var cfg Config
  if err := v.Unmarshal(&cfg); err != nil {
    return nil, err
  }

  return &cfg, nil
}

