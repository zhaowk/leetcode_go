package leetcode_go

import (
	"reflect"
	"testing"
)

func Test_majorityElement2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{nums:[]int{3,2,3}},
			want: []int{3},
		},
		{
			name: "2",
			args: args{nums:[]int{1}},
			want: []int{1},
		},
		{
			name: "3",
			args: args{nums:[]int{1,1,1,3,3,2,2,2}},
			want: []int{1,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}
