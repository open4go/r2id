package r2id

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"strconv"
	"sync"
)

const (
	DefaultB6 = 999999
	B6Bit     = 6
	B3Bit     = 3
)

var (
	node *snowflake.Node
	once sync.Once
)

func InitNode() {
	nodeId := viper.GetInt64("random.r2id.node")
	if nodeId == 0 {
		nodeId = 1
	}
	var err error
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		panic(err) // 如果节点初始化失败，则立即中止程序
	}
}

// Cutter digit with 6 bits
func Cutter(prefix, bit int) int {
	// 确保node只被初始化一次
	once.Do(InitNode)

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
	return fmt.Sprintf("%d", Cutter(3, B3Bit))
}
