package leetcode_go

var partitionResult [][]string

func partitions(s string) [][]string {

	partitionResult = make([][]string, 0)
	if len(s) == 0 {
		return partitionResult
	}

	t := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		t[i] = make([]bool, len(s)+1)
		t[i][i] = false
		for j := i + 1; j <= len(s); j++ {
			t[i][j] = isV(s[i:j])
		}
	}

	pa(s, t, 0, []string{})
	return partitionResult
}

func pa(s string, t [][]bool, start int, p []string) {

	if start == len(s) {
		tmp := make([]string, len(p))
		copy(tmp, p)
		partitionResult = append(partitionResult, tmp)
		return
	}

	for i := start + 1; i <= len(s); i++ {
		if t[start][i] {
			pa(s, t, i, append(p, s[start:i]))
		}
	}
}

func isV(s string) bool {
	if len(s) < 2 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
