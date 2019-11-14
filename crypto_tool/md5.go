package crypto_tool

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5(s string) string {

	h := md5.New()

	h.Write([]byte(s))

	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)

}

func MD5(s string) string {
	return strings.ToUpper(Md5(s))
}
