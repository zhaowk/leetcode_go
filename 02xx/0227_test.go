package _2xx

import "testing"

func Test_calculate2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "1",
			args: args{s: " 3+5 / 2 "},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate2(tt.args.s); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
