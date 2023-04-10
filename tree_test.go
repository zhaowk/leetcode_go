package _0xx

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		nodes []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{nodes: []interface{}{1, 2, 3, 4, 5}},
			want: "[1,2,3,4,5,null,null]",
		},
		{
			name: "test 2",
			args: args{nodes: []interface{}{1, nil, 3, 4, 5}},
			want: "[1,null,3,4,5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := BuildTree(tt.args.nodes...); !reflect.DeepEqual(gotRoot.String(), tt.want) {
				t.Errorf("BuildTree() = %v, want %v", gotRoot, tt.want)
			}
		})
	}
}

func TestTreeEq(t *testing.T) {
	type args struct {
		i *TreeNode
		j *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{i: BuildTree(1, 2, 3, 4, 5), j: BuildTree(1, 2, 3, 4, 5)},
			want: true,
		},
		{
			name: "test 1",
			args: args{i: BuildTree(1, 2, 3, 4, 5), j: BuildTree(1, 2, 3, 4)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeEq(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("TreeEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
