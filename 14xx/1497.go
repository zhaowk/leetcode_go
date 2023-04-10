package _4xx

func canArrange(arr []int, k int) bool {
	m := make(map[int]int, 0)

	for _, v := range arr {
		m[(k+v%k)%k]++
	}

	for i, v := range m {
		if i == 0 {
			continue
		} else if v != m[k-i] {
			return false
		}
	}
	return m[0]%2 == 0
}
