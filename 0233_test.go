package leetcode_go

import "testing"

func Test_countDigitOne(t *testing.T) {
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
			args: args{n: 0},
			want: 0,
		},
		{
			name: "max",
			args: args{n: 1000000000},
			want: 900000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
