package _0xx

func isValid(s string) bool {

	stack := make([]int32, 0)

	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, '(')
		case ')':
			if len(stack) <= 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '[':
			stack = append(stack, '[')
		case ']':
			if len(stack) <= 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '{':
			stack = append(stack, '{')
		case '}':
			if len(stack) <= 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
