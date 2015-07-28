package cryptoutils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ComputeHmac256(message string) string {
	key := []byte("big awesome secret stands here")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Hash(s string, level int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), level)
	if err != nil {
		return "", nil
	}
	return string(hashed), nil
}