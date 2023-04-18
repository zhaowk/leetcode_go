package _0xx

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zhaowk/leetcode_go"
	"testing"
)

func Test_maxAncestorDiff(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{BuildTree(8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13)},
			want: 7,
		},
		{
			name: "case 2",
			args: args{BuildTree(1, nil, 2, nil, 0, 3)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxAncestorDiff(tt.args.root), "maxAncestorDiff(%v)", tt.args.root)
		})
	}
}
