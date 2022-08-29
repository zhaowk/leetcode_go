package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNestedIterator_findNext(t *testing.T) {
	type args struct {
		nestedList []*NestedInteger
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args {[]*NestedInteger{
				{
					isInt: false,
					list: []*NestedInteger{
						{
							isInt: true,
							val: 1,
						},
						{
							isInt: true,
							val: 1,
						},
					},
				},
				{
					isInt: true,
					val: 2,
				},
				{
					isInt: false,
					list: []*NestedInteger{
						{
							isInt: true,
							val: 1,
						},
						{
							isInt: true,
							val: 1,
						},
					},
				},
			}},
			want: []int{1,1,2,1,1},
		},
		{
			name: "2",
			args: args {[]*NestedInteger{
				{
					isInt: true,
					val: 1,
				},
				{
					isInt: false,
					list: []*NestedInteger{
						{
							isInt: true,
							val: 4,
						},
						{
							isInt: false,
							list: []*NestedInteger{
								{
									isInt: true,
									val: 6,
								},
							},
						},
					},
				},
			}},
			want: []int{1,4,6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor5(tt.args.nestedList)
			res := make([]int, 0)
			for this.HasNext() {
				res = append(res, this.Next())
			}
			assert.Equal(t, tt.want, res)
		})
	}
}
