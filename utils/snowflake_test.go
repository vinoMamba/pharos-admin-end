package utils

import "testing"

func TestSnowflake(t *testing.T) {
	id := SnowNode.Generate().String()
	id2 := SnowNode.Generate().String()
	id3 := SnowNode.Generate().String()
	t.Log(id)
	t.Log(id2)
	t.Log(id3)
}
