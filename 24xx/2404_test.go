package _4xx

import "testing"

func Test_mostFrequentEven(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{0, 1, 2, 2, 4, 4, 1}}, 2},
		{"case 2", args{[]int{29, 47, 21, 41, 13, 37, 25, 7}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentEven(tt.args.nums); got != tt.want {
				t.Errorf("mostFrequentEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
