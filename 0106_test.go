package leetcode_go

import (
	"fmt"
	"testing"
)

func TestBuildTreeIP(t *testing.T) {
	fmt.Println(buildTreeIP([]int{1, 2}, []int{2, 1}))
	fmt.Println(buildTreeIP([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
