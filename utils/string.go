package utils

import (
	"math/rand"
	"time"
)

const Charset1 = "0123456789"
const Charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const Charset3 = "abcdefghijklmnopqrstuvwxyz0123456789"
const Charset4 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

func StringWithCharsetV1(length int) string {
	return StringWithCharset(length, Charset1)
}

func StringWithCharsetV2(length int) string {
	return StringWithCharset(length, Charset2)
}

func StringWithCharsetV3(length int) string {
	return StringWithCharset(length, Charset3)
}
func StringWithCharsetV4(length int) string {
	return StringWithCharset(length, Charset3)
}
