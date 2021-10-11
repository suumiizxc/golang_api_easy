package base64Upload

import (
      "os"
      "encoding/base64"
)

func Upload(fileName string, content string) error {
      decode, err := base64.StdEncoding.DecodeString(content)
      file, err := os.Create(fileName)
      defer file.Close()
      _, err = file.Write(decode)
      return err
}
