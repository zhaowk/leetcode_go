package _2xx

import "testing"
import (
	. "github.com/zhaowk/leetcode_go"
)

func Test_kthSmallest1(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{root: BuildTree(3, 1, 4, nil, 2), k: 1},
			want: 1,
		},
		{
			name: "2",
			args: args{root: BuildTree(5, 3, 6, 2, 4, nil, nil, 1), k: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest1(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest1() = %v, want %v", got, tt.want)
			}
		})
	}
}
