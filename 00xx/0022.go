package _0xx

func generateParenthesis(n int) []string {

	result := make([]string, 0)

	genParenthesis(n, 0, 0, &result, "")

	return result
}

func genParenthesis(n, l, r int, res *[]string, pi string) {

	if n == l && n == r {
		*res = append(*res, pi)
	}

	if l <= n {
		genParenthesis(n, l+1, r, res, pi+"(")
	}

	if r <= n && r < l {
		genParenthesis(n, l, r+1, res, pi+")")
	}
}
