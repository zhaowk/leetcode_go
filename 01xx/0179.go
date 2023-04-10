package _1xx

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {

	sort.Slice(nums, func(i, j int) bool {
		s1 := strconv.Itoa(nums[i])
		s2 := strconv.Itoa(nums[j])

		for {
			if s1 == s2 {
				return true // 相等
			}

			if len(s1) == 0 {
				s1 = strconv.Itoa(nums[i])
			}

			if len(s2) == 0 {
				s2 = strconv.Itoa(nums[j])
			}

			if s1[0] > s2[0] {
				return true
			} else if s1[0] < s2[0] {
				return false
			}
			s1, s2 = s1[1:], s2[1:]
		}
	})

	res := ""
	for _, v := range nums {
		res += strconv.Itoa(v)
	}

	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}

	return res
}
