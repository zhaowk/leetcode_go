package _2xx

func findOrder(numCourses int, prerequisites [][]int) []int {
	// BFS
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		q      = make([]int, 0)
		result = make([]int, 0)
	)

	for _, val := range prerequisites {
		edges[val[1]] = append(edges[val[1]], val[0])
		indeg[val[0]]++
	}

	for idx, val := range indeg {
		if val == 0 {
			q = append(q, idx)
		}
	}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		result = append(result, item)

		for _, val := range edges[item] {
			indeg[val]--
			if indeg[val] == 0 {
				q = append(q, val)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}

	return []int{}
}

func findOrderDFS(numCourses int, prerequisites [][]int) []int {
	return []int{}
}
