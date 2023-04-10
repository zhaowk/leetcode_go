package _7xx

func isBipartite(graph [][]int) bool {
	bi1 := make(map[int]bool)
	bi2 := make(map[int]bool)

	vi := make(map[int]bool)

	if len(graph) < 2 {
		return true
	}

	bi1[0] = true

	for len(vi) < len(graph) {
		bl1 := len(bi1)
		bl2 := len(bi2)
		for k := range bi1 {
			if !vi[k] {
				vi[k] = true
				for _, n := range graph[k] {
					if bi1[n] {
						return false
					}
					bi2[n] = true
					//vi[n] = true
				}
			}
		}

		for k := range bi2 {
			if !vi[k] {
				vi[k] = true
				for _, n := range graph[k] {
					if bi2[n] {
						return false
					}
					bi1[n] = true
					//vi[n] = true
				}
			}
		}

		if bl1 == len(bi1) && bl2 == len(bi2) {
			for i := 0; i < len(graph); i++ {
				if !vi[i] {
					bi1[i] = true
					break
				}
			}
		}
	}

	return true
}
