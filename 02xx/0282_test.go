package _2xx

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "123 = 6",
			args: args{num: "123", target: 6},
			want: []string{"1+2+3", "1*2*3"},
		},
		{
			name: "232 = 8",
			args: args{num: "232", target: 8},
			want: []string{"2+3*2", "2*3+2"},
		},
		{
			name: "105 = 5",
			args: args{num: "105", target: 5},
			want: []string{"1*0+5", "10-5"},
		},
		{
			name: "3456237490 = 9191",
			args: args{num: "3456237490", target: 9191},
			want: []string{},
		},
		{
			name: "00 = 0",
			args: args{num: "00", target: 0},
			want: []string{"0+0", "0-0", "0*0"},
		},
		{
			name: "0 = 0",
			args: args{num: "0", target: 0},
			want: []string{"0"},
		},
		{
			name: "01 = 1",
			args: args{num: "01", target: 1},
			want: []string{"0+1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
