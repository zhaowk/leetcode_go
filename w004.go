package leetcode_go

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {

	noDepend := make([]int, n+1)
	depend := make(map[int][]int, 0)

	for i := 1; i <= n; i++ {
		noDepend[i] = i
	}

	for _, v := range dependencies {
		depend[v[1]] = append(depend[v[1]], v[0])
		noDepend[v[0]] = 0
		noDepend[v[1]] = 0
	}

	tmp := make([]int, 0)
	for _, v := range noDepend {
		if v != 0 {
			tmp = append(tmp, v)
		}
	}

	picked := 0
	count := 0

	for picked <= n {

		//picked = pick(k,)
		count++
	}

	return count
}
