package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	type args struct {
		obj MedianFinder
		op []struct {
			typ string // add, find
			num int // add an integer
			res float64  // find median result
		}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "basic",
			args: args{
				obj: Constructor3(),
				op: []struct {
					typ string // add, find
					num int // add an integer
					res float64  // find median result
				}{
					{
						typ: "add",
						num: 1,
					},
					{
						typ: "add",
						num: 2,
					},
					{
						typ: "find",
						res: 1.5,
					},
					{
						typ: "add",
						num: 3,
					},{
						typ: "find",
						res: 2,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.op {
				switch v.typ {
				case "add":
					tt.args.obj.AddNum(v.num)
				case "find":
					assert.Equal(t, v.res, tt.args.obj.FindMedian())
				default:
					t.Fatal("op not found")
				}
			}
		})
	}
}
