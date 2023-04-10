package _4xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinStep(t *testing.T) {
	type args struct {
		board string
		hand  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{board: "WRRBBW", hand: "RB"},
			want: -1,
		},
		{
			name: "2",
			args: args{board: "WWRRBBWW", hand: "WRBRW"},
			want: 2,
		},
		{
			name: "3",
			args: args{board: "G", hand: "GGGGG"},
			want: 2,
		},
		{
			name: "4",
			args: args{board: "RBYYBBRRB", hand: "YRBGB"},
			want: 3,
		},
		{
			name: "5",
			args: args{board: "RRWWRRBBRR", hand: "WB"},
			want: 2,
		},
		{
			name: "6",
			args: args{board: "BRWGWYY", hand: "YGBWY"},
			want: -1,
		},
		// TODO
		//{
		//	name: "7",
		//	args: args{board: "RRYGGYYRRYYGGYRR", hand: "GGBBB"},
		//	want: 5,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMinStep(tt.args.board, tt.args.hand))
		})
	}
}
