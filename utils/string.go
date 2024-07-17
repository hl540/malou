package utils

import (
	"math/rand"
	"time"
)

const Charset1 = "0123456789"
const Charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const Charset3 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 随机数种子
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset 生成指定长度的随机字符串
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
