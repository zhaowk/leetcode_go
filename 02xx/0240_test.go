package _2xx

import "testing"

func Test_searchMatrix1(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
	}
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: matrix,
				target: 5,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				matrix: matrix,
				target: 20,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				matrix: matrix,
				target: 30,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				matrix: matrix,
				target: 1000,
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				matrix: matrix,
				target: 0,
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				matrix: matrix,
				target: 25,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				matrix: matrix,
				target: 10,
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				matrix: [][]int{{1, 2, 3, 4, 5}},
				target: 2,
			},
			want: true,
		},
		{
			name: "9",
			args: args{
				matrix: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
				target: 15,
			},
			want: true,
		},
		{
			name: "10",
			args: args{
				matrix: [][]int{{3, 5, 9, 9, 9, 11}, {5, 8, 13, 13, 16, 17}, {10, 10, 14, 14, 16, 19}, {15, 18, 20, 24, 26, 26}, {20, 24, 29, 32, 37, 41}},
				target: 32,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix1() = %v, want %v", got, tt.want)
			}
		})
	}
}
