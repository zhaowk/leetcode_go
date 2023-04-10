package _3xx

import "testing"

func Test_maxProduct_(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			want: 16,
		},
		{
			name: "2",
			args: args{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			want: 4,
		},
		{
			name: "3",
			args: args{[]string{"a", "aa", "aaa", "aaaa"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct_(tt.args.words); got != tt.want {
				t.Errorf("maxProduct_() = %v, want %v", got, tt.want)
			}
		})
	}
}
