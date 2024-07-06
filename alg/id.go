package alg

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	snowflake.NodeBits = 1
	snowflake.Epoch = 1681987346000
	snowflake.StepBits = 10
	n, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}
	node = n
}

// 生成id
// StepBits 10位
func Id() (int64, error) {
	id := node.Generate()
	return id.Int64(), nil
}
func Sid() (string, error) {
	id := node.Generate()
	return id.String(), nil
}
