package _0xx

import (
	"reflect"
	"testing"
)

func TestBuildList(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test build list",
			args: args{n: []int{1, 2, 3}},
			want: "[1,2,3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildList(tt.args.n...); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("BuildList() = %v, want %v", got, tt.want)
			}
		})
	}
}
