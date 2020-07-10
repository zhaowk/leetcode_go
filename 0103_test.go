package leetcode_go

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	fmt.Println(zigzagLevelOrder(buildTree(3, 9, 20, nil, nil, 15, 7)))
}
