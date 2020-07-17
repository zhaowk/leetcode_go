package leetcode_go

import "strconv"

func evalRPN(tokens []string) int {

	stack := make([]int, 0)

	for _, token := range tokens {
		switch token {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			res, err := strconv.Atoi(token)
			if err != nil {
				res = 0
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
