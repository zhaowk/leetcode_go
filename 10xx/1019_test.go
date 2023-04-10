package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				head: BuildList(2, 1, 5),
			},
			want: []int{5, 5, 0},
		},
		{
			name: "test 2",
			args: args{
				head: BuildList(2, 7, 4, 3, 5),
			},
			want: []int{7, 0, 5, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextLargerNodes(tt.args.head), "nextLargerNodes(%v)", tt.args.head)
		})
	}
}
