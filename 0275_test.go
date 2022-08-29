package leetcode_go

import "testing"

func Test_hIndex2(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{citations: []int{0,1,3,5,6}},
			want: 3,
		},
		{
			name: "2",
			args: args{citations: []int{0,0,1}},
			want: 1,
		},
		{
			name: "3",
			args: args{citations: []int{0,0,0}},
			want: 0,
		},
		{
			name: "4",
			args: args{citations: []int{0,0}},
			want: 0,
		},
		{
			name: "5",
			args: args{citations: []int{0,1}},
			want: 1,
		},
		{
			name: "6",
			args: args{citations: []int{0,0,0,0,0}},
			want: 0,
		},
		{
			name: "7",
			args: args{citations: []int{0,0,4,4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex2(tt.args.citations); got != tt.want {
				t.Errorf("hIndex2() = %v, want %v", got, tt.want)
			}
		})
	}
}
