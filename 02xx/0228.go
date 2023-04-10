package _2xx

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	start, end := nums[0], nums[0]
	res := make([]string, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			end = nums[i-1]
			if start == end {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
		}
		end = nums[i]
	}

	if start == end {
		res = append(res, fmt.Sprintf("%d", start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, end))
	}

	return res
}
