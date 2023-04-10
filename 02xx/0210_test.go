package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOrder(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		result        []int
	}{
		{
			4,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[]int{0, 1, 2, 3},
		},
		{
			2,
			[][]int{{1, 0}},
			[]int{0, 1},
		},
		{
			2,
			[][]int{},
			[]int{0, 1},
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.result, findOrder(v.numCourses, v.prerequisites))
	}
}
