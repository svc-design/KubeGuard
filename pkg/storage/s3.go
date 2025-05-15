func UploadFile(sess *session.Session, bucket, key, filePath string) error {
  svc := s3.New(sess)
  file, err := os.Open(filePath)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = svc.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
    Body:   file,
  })
  return err
}

func DownloadFile(sess *session.Session, bucket, key, destPath string) error {
  svc := s3.New(sess)
  obj, err := svc.GetObject(&s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
  })
  if err != nil {
    return err
  }
  defer obj.Body.Close()

  file, err := os.Create(destPath)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.Copy(file, obj.Body)
  return err
}

