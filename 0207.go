package leetcode_go

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	queue := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		adjacency[i] = make([]int, 0)
	}

	// Get the indegree and adjacency of every course.
	for _, cp := range prerequisites {
		indegrees[cp[0]]++
		adjacency[cp[1]] = append(adjacency[cp[1]], cp[0])
	}
	// Get all the courses with the indegree of 0.
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS TopSort.
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		for _, cur := range adjacency[pre] {
			indegrees[cur]--
			if indegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}
	return numCourses == 0
}
