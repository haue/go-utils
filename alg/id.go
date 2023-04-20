package alg

import (
	"github.com/bwmarrin/snowflake"
)

// 生成id
// StepBits 10位
func IdGen() (int64, error) {
	snowflake.NodeBits = 1
	snowflake.Epoch = 1681987346000
	snowflake.StepBits = 10
	node, err := snowflake.NewNode(0)
	if err != nil {
		return 0, err
	}
	id := node.Generate()
	return id.Int64(), nil
}
