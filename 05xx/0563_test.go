package _5xx

import "testing"
import (
	. "github.com/zhaowk/leetcode_go"
)

func Test_findTilt(t *testing.T) {
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
			args: args{root: BuildTree(1, 2, 3)},
			want: 1,
		},
		{
			name: "2",
			args: args{root: BuildTree(4, 2, 9, 3, 5, nil, 7)},
			want: 15,
		},
		{
			name: "3",
			args: args{root: BuildTree(21, 7, 14, 1, 1, 2, 2, 3, 3)},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
