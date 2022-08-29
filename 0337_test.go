package leetcode_go

import "testing"

func Test_rob2(t *testing.T) {
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
			args: args{buildTree(3,2,3,nil,3,nil,1)},
			want: 7,
		},
		{
			name: "2",
			args: args{buildTree(3,4,5,1,3,nil,1)},
			want: 9,
		},
		{
			name: "3",
			args: args{buildTree(79,99,77,nil,nil,nil,69,nil,60,53,nil,73,11,nil,nil,nil,62,27,62,nil,nil,98,50,nil,nil,90,48,82,nil,nil,nil,55,64,nil,nil,73,56,6,47,nil,93,nil,nil,75,44,30,82,nil,nil,nil,nil,nil,nil,57,36,89,42,nil,nil,76,10,nil,nil,nil,nil,nil,32,4,18,nil,nil,1,7,nil,nil,42,64,nil,nil,39,76,nil,nil,6,nil,66,8,96,91,38,38,nil,nil,nil,nil,74,42,nil,nil,nil,10,40,5,nil,nil,nil,nil,28,8,24,47,nil,nil,nil,17,36,50,19,63,33,89,nil,nil,nil,nil,nil,nil,nil,nil,94,72,nil,nil,79,25,nil,nil,51,nil,70,84,43,nil,64,35,nil,nil,nil,nil,40,78,nil,nil,35,42,98,96,nil,nil,82,26,nil,nil,nil,nil,48,91,nil,nil,35,93,86,42,nil,nil,nil,nil,0,61,nil,nil,67,nil,53,48,nil,nil,82,30,nil,97,nil,nil,nil,1,nil,nil)},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.root); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}
