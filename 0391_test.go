package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isRectangleCover(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				rectangles: [][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				rectangles: [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}},
			},
			want: false,
		},
		// [[0,0,4,1},{7,0,8,2},{6,2,8,3},{5,1,6,3},{4,0,5,1},{6,0,7,2},{4,2,5,3},{2,1,4,3},{0,1,2,2},{0,2,2,3},{4,1,5,2},{5,0,6,1]]
		{
			name: "5",
			args: args{
				rectangles: [][]int{{0, 0, 4, 1}, {7, 0, 8, 2}, {6, 2, 8, 3}, {5, 1, 6, 3}, {4, 0, 5, 1}, {6, 0, 7, 2}, {4, 2, 5, 3}, {2, 1, 4, 3}, {0, 1, 2, 2}, {0, 2, 2, 3}, {4, 1, 5, 2}, {5, 0, 6, 1}},
			},
			want: true,
		},
		// [[0,0,1,1},{0,1,3,2},{1,0,2,2]]
		{
			name: "6",
			args: args{
				rectangles: [][]int{{0, 0, 1, 1}, {0, 1, 3, 2}, {1, 0, 2, 2}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isRectangleCover(tt.args.rectangles))
		})
	}
}
