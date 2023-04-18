package _4xx

import "testing"

func Test_countDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{"08-15", "08-18", "08-16", "08-19"},
			3,
		},
		{
			"case 2",
			args{"10-01", "10-31", "11-01", "12-31"},
			0,
		},
		{
			"case 3",
			args{"01-15", "10-28", "08-05", "12-31"},
			85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob); got != tt.want {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
}
