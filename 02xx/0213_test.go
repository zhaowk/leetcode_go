package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{nums: []int{2, 3, 2}},
			want: 3,
		},
		{
			name: "basic 2",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "basic 3",
			args: args{nums: []int{0}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rob(tt.args.nums))
		})
	}
}
