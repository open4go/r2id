package r2id

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
)

const (
	DefaultB6 = 999999
	B6Bit     = 6
	B3Bit     = 3
)

// Cutter digit with 6 bits
func Cutter(prefix, bit int) int {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return DefaultB6
	}

	// Generate a snowflake ID.
	id := node.Generate()
	s := fmt.Sprintf("%d", id)
	if len(s) > bit {
		s = s[len(s)-bit:]
	}

	// 加前缀
	s = fmt.Sprintf("%d%s", prefix, s)

	d, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return DefaultB6
	}
	return d
}

func D6Code() int {
	return Cutter(6, B6Bit)
}

func D3Code() int {
	return Cutter(3, B3Bit)
}

func S6Code() string {
	return fmt.Sprintf("%d", Cutter(6, B6Bit))
}

func S3Code() string {
	return fmt.Sprintf("%d", Cutter(3, B6Bit))
}
