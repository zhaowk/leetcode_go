package _9xx

import "testing"

func Test_distinctSubseqII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{""}, 0},
		{"case 2", args{"a"}, 1},
		{"case 3", args{"aaa"}, 3},
		{"case 4", args{"blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr"}, 589192369},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctSubseqII(tt.args.s); got != tt.want {
				t.Errorf("distinctSubseqII() = %v, want %v", got, tt.want)
			}
		})
	}
}
