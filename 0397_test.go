package leetcode_go

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{n:7},
			want: 4,
		},
		{
			name: "2",
			args: args{n:8},
			want: 3,
		},
		{
			name: "3",
			args: args{n:4},
			want: 2,
		},
		{
			name: "4",
			args: args{n:15},
			want: 5,
		},
		{
			name: "5",
			args: args{n:1500000000},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
