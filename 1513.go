package leetcode_go

func numSub(s string) int {

	res := 0

	idx := 0
	for idx < len(s) {
		j := idx
		for j < len(s) && s[j] == '1' {
			j++
		}

		if j > idx {
			n := j - idx
			if n%2 == 1 {
				res = (res + (n+1)/2*n) % 1000000007
			} else {
				res = (res + n/2*(n+1)) % 1000000007
			}
			idx = j
		} else {
			idx++
		}
	}

	return res
}
