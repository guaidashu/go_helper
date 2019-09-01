package crypto_tool

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}
