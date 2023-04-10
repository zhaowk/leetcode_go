package _1xx

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	if len(s) == len(t) {
		if s != t {
			return 0
		} else {
			return 1
		}
	}

	tmp := make([]int, len(s)+1)
	for j := 0; j <= len(s); j++ {
		tmp[j] = 1
	}

	for i := 1; i <= len(t); i++ {

		prev := tmp[0]
		tmp[0] = 0

		for j := 1; j <= len(s); j++ {

			if s[j-1] == t[i-1] {
				prev, tmp[j] = tmp[j], tmp[j-1]+prev
			} else {
				prev, tmp[j] = tmp[j], tmp[j-1]
			}
		}
	}

	return tmp[len(s)]
}

// 超时
func numDistinct1(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	if len(s) == len(t) {
		if s != t {
			return 0
		} else {
			return 1
		}
	}

	tmp := make([]int, len(t))
	res := 0
	numDistincts(s, t, tmp, 0, &res)
	return res
}

func numDistincts(s string, t string, tmp []int, idx int, r *int) {
	if idx < len(t) {
		i := 0
		if idx > 1 {
			i = tmp[idx-1]
		}
		for ; i < len(s); i++ {
			if s[i] == t[idx] {
				tmp[idx] = i
				if idx == 0 || tmp[idx] > tmp[idx-1] {
					numDistincts(s, t, tmp, idx+1, r)
				}
			}
		}
	} else {
		*r++
		return
	}
}
