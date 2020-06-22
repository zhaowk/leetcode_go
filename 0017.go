package leetcode_go

var s = map[int]string{2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz"}

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	for _, digit := range digits {
		if len(result) == 0 {
			for _, letter := range s[int(digit-'0')] {
				result = append(result, string(letter))
			}
		} else {
			midResult := make([]string, 0)
			for _, r := range result {
				for _, letter := range s[int(digit-'0')] {
					midResult = append(midResult, r+string(letter))
				}
			}
			result = midResult
		}
	}

	return result
}
