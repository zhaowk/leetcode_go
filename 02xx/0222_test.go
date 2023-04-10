package _2xx

import "testing"
import (
	. "github.com/zhaowk/leetcode_go"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{root: BuildTree()},
			want: 0,
		},
		{
			name: "2",
			args: args{root: BuildTree(1)},
			want: 1,
		},
		{
			name: "3",
			args: args{root: BuildTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
