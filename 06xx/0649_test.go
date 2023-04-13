package _6xx

import "testing"

func Test_predictPartyVictory(t *testing.T) {
	type args struct {
		senate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 1", args{"RD"}, "Radiant"},
		{"case 2", args{"RDD"}, "Dire"},
		{"case 3", args{"DDRRR"}, "Dire"},
		{"case 4", args{"RRDDD"}, "Radiant"},
		{"case 5", args{"DRRDRDRDRDDRDRDR"}, "Radiant"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPartyVictory(tt.args.senate); got != tt.want {
				t.Errorf("predictPartyVictory() = %v, want %v", got, tt.want)
			}
		})
	}
}
