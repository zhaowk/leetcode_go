package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumOfBeauties(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 2, 3}}, 2},
		{"case 2", args{[]int{2, 4, 6, 4}}, 1},
		{"case 3", args{[]int{3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumOfBeauties(tt.args.nums), "sumOfBeauties(%v)", tt.args.nums)
		})
	}
}
