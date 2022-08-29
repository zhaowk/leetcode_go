package leetcode_go

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	t1 := buildTree(3,5,1,6,2,0,8,nil,nil,7,4)
	t2 := buildTree(1, 2)
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{root: t1, p: t1.Left, q: t1.Left.Right.Right},
			want: t1.Left,
		},
		{
			name: "2",
			args: args{root: t2, p: t2, q: t2.Left},
			want: t2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
