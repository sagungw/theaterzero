package encryptor

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
)

const defaultByteSize = 8

// GenerateSalt :nodoc:
func GenerateSalt() (result string, err error) {
	b := make([]byte, defaultByteSize)
	_, err = rand.Read(b)
	if err != nil {
		return
	}

	result = base64.URLEncoding.EncodeToString(b)
	return
}

// EncryptWithSalt :nodoc:
func EncryptWithSalt(s string, salt string) []byte {
	b := sha1.Sum([]byte(s + salt))
	return b[:]
}
