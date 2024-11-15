package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func EncryPassword(password string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
