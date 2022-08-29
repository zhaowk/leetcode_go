package leetcode_go

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{1,2,3,5}, k:1},
			want: []int{1,5},
		},
		{
			name: "2",
			args: args{arr: []int{1,2,3,5}, k:2},
			want: []int{1,3},
		},
		{
			name: "3",
			args: args{arr: []int{1,2,3,5}, k:3},
			want: []int{2,5},
		},
		{
			name: "4",
			args: args{arr: []int{1,2,3,5}, k:4},
			want: []int{1,2},
		},
		{
			name: "5",
			args: args{arr: []int{1,2,3,5}, k:5},
			want: []int{3,5},
		},
		{
			name: "6",
			args: args{arr: []int{1,2,3,5}, k:6},
			want: []int{2,3},
		},
		{
			name: "7",
			args: args{arr: []int{1,7}, k:1},
			want: []int{1,7},
		},
		{
			name: "8",
			args: args{arr: []int{1,1553,3491,3803,4153,4561,4643,4723,4813,7411,8707,9511,9791,11719,18269,21493,21961,23833,25303,26717}, k:58},
			want: []int{4561,21961},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
