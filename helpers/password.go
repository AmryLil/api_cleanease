package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

type hash struct{}

type HashInterface interface {
	HashPassword(password string) string
	CompareHash(password, hashed string) bool
}

func NewHash() HashInterface {
	return &hash{}
}

func (h hash) HashPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (h hash) CompareHash(password, hashed string) bool {
	return h.HashPassword(password) == hashed
}
