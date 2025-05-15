var Logger = logrus.New()

func init() {
  Logger.SetFormatter(&logrus.TextFormatter{
    FullTimestamp: true,
  })
  Logger.SetLevel(logrus.InfoLevel)
}

