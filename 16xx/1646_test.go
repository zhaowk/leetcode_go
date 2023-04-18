package _6xx

import "testing"

func Test_getMaximumGenerated(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{7}, 3},
		{"case 2", args{100}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGenerated(tt.args.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
