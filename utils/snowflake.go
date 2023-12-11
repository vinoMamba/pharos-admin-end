package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var SnowNode *snowflake.Node

func init() {
	var err error
	SnowNode, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
}

func GetSnowflakeIdInt64() int64 {
	return SnowNode.Generate().Int64()
}
