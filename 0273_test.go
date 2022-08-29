package leetcode_go

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1000",
			args: args{num: 1000},
			want: "One Thousand",
		},
		{
			name: "20000000",
			args: args{num: 20000000},
			want: "Twenty Million",
		},
		{
			name: "1520000000",
			args: args{num: 1520000000},
			want: "One Billion Five Hundred Twenty Million",
		},
		{
			name: "1234567",
			args: args{num: 1234567},
			want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name: "123456789",
			args: args{num: 123456789},
			want: "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
