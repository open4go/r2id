package r2id

import (
	"context"
	"fmt"
	"github.com/r2day/db"
	"log"
	"time"
)

const (
	prefix = "open4go:r2id:ticket:"
)

// QueueTicketNumber 获取排队号自动增加
// 每天自动过期
func QueueTicketNumber(ctx context.Context, ticketPrefix string) string {
	// 获取当前日期作为 Redis Key
	key := time.Now().Format("2006-01-02")
	fullKey := fmt.Sprintf("%s:%s", prefix, key)

	// 获取当前排队号
	currentNumber, err := db.RDB.Incr(ctx, fullKey).Result()
	if err != nil {
		log.Fatal(err)
	}

	// 两天后过期
	db.RDB.Expire(ctx, fullKey, time.Hour*48)

	// 格式化排队号，例如 A5001
	formattedNumber := fmt.Sprintf("%s%04d", ticketPrefix, currentNumber)
	return formattedNumber

}
