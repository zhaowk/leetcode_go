package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_camelMatch(t *testing.T) {
	type args struct {
		queries []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			"case 1",
			args{[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"},
			[]bool{true, false, true, true, false},
		},
		{
			"case 2",
			args{[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"},
			[]bool{true, false, true, false, false},
		},
		{
			"case 3",
			args{[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"},
			[]bool{false, true, false, false, false},
		},
		{
			"case 4",
			args{[]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "CooP"},
			[]bool{false, false, true},
		},
		{
			"case 5",
			args{[]string{"uAxaqlzahfialcezsLfj", "cAqlzyahaslccezssLfj", "AqlezahjarflcezshLfj", "AqlzofahaplcejuzsLfj", "tAqlzahavslcezsLwzfj", "AqlzahalcerrzsLpfonj", "AqlzahalceaczdsosLfj", "eAqlzbxahalcezelsLfj"}, "AqlzahalcezsLfj"},
			[]bool{true, true, true, true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, camelMatch(tt.args.queries, tt.args.pattern), "camelMatch(%v, %v)", tt.args.queries, tt.args.pattern)
		})
	}
}
