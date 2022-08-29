package leetcode_go

const (
	opAdd = iota
	opMinus
	opMultiply
)

var (
	opM = map[int]func(i, j int) int {
		opAdd: func(i, j int) int {return i + j},
		opMinus: func(i, j int) int {return i - j},
		opMultiply: func(i, j int) int {return i * j},
	}
)

func diffWaysToCompute(expression string) []int {
	nums := make([]int, 0)
	operator := make([]int, 0)

	num := 0
	isNum := true

	for _, e := range expression {
		switch e {
		case '+':
			if isNum {
				nums = append(nums, num)
			}
			operator = append(operator, opAdd)
			isNum = false
		case '-':
			if isNum {
				nums = append(nums, num)
			}
			operator = append(operator, opMinus)
			isNum = false
		case '*':
			if isNum {
				nums = append(nums, num)
			}
			operator = append(operator, opMultiply)
			isNum = false
		case '0','1','2','3','4','5','6','7','8','9':
			isNum = true
			num = 10 * num + int(e - '0')
		default:
			if isNum {
				nums = append(nums, num)
			}
			isNum = false
		}

		if !isNum {
			num = 0
		}
	}

	if isNum {
		nums = append(nums, num)
	}

	var compute func(ops, ns []int) []int

	compute = func(ops, ns []int) []int {
		if len(ops) == 0 && len(ns) == 1 {
			return []int{ns[0]}
		} else if len(ops) == 1 && len(ns) == 2 {
			return []int{opM[ops[0]](ns[0], ns[1])}
		}

		res := make([]int, 0)

		for i, v := range ops {
			left := compute(ops[:i], ns[:i + 1])
			var right []int
			if i != len(ops) - 1 {
				right = compute(ops[i + 1:], ns[i+1:])
			} else {
				right = []int{ns[len(ns) - 1]}
			}

			for _, l := range left {
				for _, r := range right {
					res  = append(res, opM[v](l, r))
				}
			}
		}

		return res
	}

	return compute(operator, nums)
}
