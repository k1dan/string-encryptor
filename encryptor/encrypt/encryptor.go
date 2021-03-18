package encrypt

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncryptString(strings string) (encryptedString string) {
	h := sha256.New()
	h.Write([]byte(strings))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}
