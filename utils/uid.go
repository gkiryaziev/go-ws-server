package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

func generateRandomBytes(size int) []byte {
	buffer := make([]byte, size)
	rand.Read(buffer)
	return buffer
}

// generate random id as md5 string with specified size
func GenerateRandomMD5String(size int) string {
	hash := md5.New()
	hash.Write(generateRandomBytes(size))
	return hex.EncodeToString(hash.Sum(nil))
}
