package _1xx

import "sort"
import (
	. "github.com/zhaowk/leetcode_go"
)

func maxPoints(points [][]int) int {

	if len(points) < 3 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	mx := 2

	xy := make(map[int]map[int]int)

	for _, point := range points {
		if xy[point[0]] == nil {
			xy[point[0]] = map[int]int{point[1]: 1}
		} else {
			xy[point[0]][point[1]]++
		}
	}

	for i, point := range points {
		x1, y1 := point[0], point[1]

		l := 0
		for _, v := range xy[x1] {
			l += v
		}

		mx = MyMax(mx, l)
		for j := i + 1; j < len(points); j++ {
			if points[j][0] == x1 { //
				continue
			}
			x2, y2 := points[j][0], points[j][1]
			a, b, c := y1-y2, x1*y2-x2*y1, x1-x2
			l := xy[x1][y1] + xy[x2][y2]
			for k := j + 1; k < len(points); k++ {
				if points[k][0] == x2 {
					continue
				}
				x := points[k][0]
				cy := a*x + b
				if cy%c == 0 && xy[x] != nil && xy[x][cy/c] > 0 {
					l += xy[x][cy/c]
					for k+1 < len(points) && points[k+1][0] == x {
						k++
					}
				}
			}
			mx = MyMax(mx, l)
		}
	}

	return mx

}
