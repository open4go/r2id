package r2id

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

const (
	OrderLength = 18
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

func GenerateOutTradeNo(prefix ...string) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder

	// 如果用户提供了前缀，使用第一个前缀，否则默认前缀为 "A"
	if len(prefix) > 0 {
		result.WriteString(prefix[0])
	} else {
		result.WriteString("A")
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}

	// 获取当前时间，并使用时间作为订单号的一部分
	dt := time.Now().In(loc)
	result.WriteString(dt.Format("20060102150405"))

	// 随机生成订单号的剩余部分，确保总长度为 32
	for i := 0; i < OrderLength-len(result.String()); i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result.WriteByte(charset[n.Int64()])
	}

	return result.String()
}
