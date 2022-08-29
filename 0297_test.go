package leetcode_go

import (
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{root: buildTree(1,2,3,nil,nil,4,5)},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			got := this.serialize(tt.args.root);


			node := this.deserialize(got)

			fmt.Println(node)

			fmt.Println(treeEq(tt.args.root, node))

		})
	}
}
