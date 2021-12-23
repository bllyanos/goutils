package strings

import (
	"math/rand"
)

type RandomCharset string

const (
	AlphaCharset      RandomCharset = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpperCharset RandomCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaMixCharset   RandomCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Seed(v int64) {
	rand.Seed(v)
}

func RandomString(charset RandomCharset, strLen int32) string {
	csetLen := len(charset)
	result := ""
	for i := 0; i < int(strLen); i++ {
		charIdx := rand.Intn(csetLen)
		char := string(charset[charIdx])
		result = result + char
	}
	return result
}
