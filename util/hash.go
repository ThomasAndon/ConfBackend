package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func calcSHA256(input []byte) string {
	hash := sha256.New()
	hash.Write(input)
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// 获取消息的SHA256的前4位
func FourDigitSHA256(input []byte) []byte {
	hash := calcSHA256(input)
	if len(hash) >= 4 {
		return []byte(hash[:4])
	}
	return nil
}
