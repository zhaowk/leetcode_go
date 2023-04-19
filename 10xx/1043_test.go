package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 15, 7, 9, 2, 5, 10}, 3}, 84},
		{"case 2", args{[]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4}, 83},
		{"case 3", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSumAfterPartitioning(tt.args.arr, tt.args.k), "maxSumAfterPartitioning(%v, %v)", tt.args.arr, tt.args.k)
		})
	}
}
