package ms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findClosedNumbers(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test4",
			args: args{
				num: 4,
			},
			want: []int{8, 2},
		},
		{
			name: "test3",
			args: args{
				num: 3,
			},
			want: []int{5, -1},
		},
		{
			name: "test2",
			args: args{
				num: 2,
			},
			want: []int{4, 1},
		},
		{
			name: "test1",
			args: args{
				num: 1,
			},
			want: []int{2, -1},
		},
		{
			name: "test67",
			args: args{
				num: 67,
			},
			want: []int{69, 56},
		},
		{
			name: "test124",
			args: args{
				num: 124,
			},
			want: []int{143, 122},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findClosedNumbers(tt.args.num), "findClosedNumbers(%v)", tt.args.num)
		})
	}
}
