package leetcode_go

import "strings"

func wordBreak1(s string, wordDict []string) []string {
	ls, lw := len(s), len(wordDict)

	if ls == 0 || lw == 0 {
		return []string{}
	}

	resMap := make([][]int, ls+1)
	resMap[0] = make([]int, 1)
	for i := 0; i < ls; i++ {
		if len(resMap[i]) > 0 {
			for _, word := range wordDict {
				if i+len(word) <= ls && s[i:i+len(word)] == word {
					resMap[i+len(word)] = append(resMap[i+len(word)], i)
				}
			}
		}
	}
	result := make([]string, 0)
	output(s, resMap, []string{}, &result)
	return result
}

func output(s string, resMap [][]int, tmp []string, res *[]string) {
	if len(s) == 0 || len(resMap) <= 1 || len(resMap[len(resMap)-1]) == 0 { // empty
		return
	}

	for _, r := range resMap[len(resMap)-1] {
		tmp = append(tmp, s[r:len(resMap)-1])
		if r == 0 {
			t := make([]string, len(tmp))
			for i := range tmp {
				t[len(t)-1-i] = tmp[i]
			}
			*res = append(*res, strings.Join(t, " "))
		}
		output(s, resMap[:r+1], tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}
