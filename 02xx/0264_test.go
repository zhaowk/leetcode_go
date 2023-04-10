package _2xx

import "testing"

func Test_nthUglyNumber(t *testing.T) {
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
			args: args{n: 1},
			want: 1,
		},
		{
			name: "3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "5",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "10",
			args: args{n: 10},
			want: 12,
		},
		{
			name: "1690",
			args: args{n: 1690},
			want: 2123366400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
