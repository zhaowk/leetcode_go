package _2xx

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{k: 3, n: 7},
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "2",
			args: args{k: 3, n: 9},
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "3",
			args: args{k: 4, n: 20},
			want: [][]int{{1, 2, 8, 9}, {1, 3, 7, 9}, {1, 4, 6, 9}, {1, 4, 7, 8}, {1, 5, 6, 8},
				{2, 3, 6, 9}, {2, 3, 7, 8}, {2, 4, 5, 9}, {2, 4, 6, 8}, {2, 5, 6, 7}, {3, 4, 5, 8}, {3, 4, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
