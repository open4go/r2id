package r2id

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result.WriteByte(charset[n.Int64()])
	}
	return result.String()
}

// GenerateOutTradeNo 生成商户订单号
func GenerateOutTradeNo() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-|*"
	var result strings.Builder
	//timestamp := time.Now().UnixNano() / int64(time.Millisecond) // 获取当前时间戳（毫秒）
	result.WriteString(time.Now().Format("20060102150405")) // 使用时间作为订单号的一部分
	for i := 0; i < 32-len(result.String()); i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result.WriteByte(charset[n.Int64()])
	}
	return result.String()
}
