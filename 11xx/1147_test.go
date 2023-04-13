package _1xx

import "testing"

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{"ghiabcdefhelloadamhelloabcdefghi"}, 7},
		{"case 2", args{"merchant"}, 1},
		{"case 3", args{"antaprezatepzapreanta"}, 11},
		{"case 4", args{"antaprezatatzapreanta"}, 13},
		{"case 5", args{"elvtoelvto"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
