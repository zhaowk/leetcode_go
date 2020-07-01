package leetcode_go

func minSubArrayLen(s int, nums []int) int {

	ln := len(nums)
	if ln == 0 {
		return 0
	}

	if nums[0] == s {
		return 1
	}

	mi := 2147483647
	start, end := 0, 0
	sum := 0
	found := false

	for end < ln {
		if nums[end] >= s {
			return 1
		} else {
			sum += nums[end]
			if sum >= s {
				found = true
				mi = min(mi, end-start+1)
				sum -= nums[start]
				sum -= nums[end]
				start++
			} else {
				end++
			}

			if start > end {
				break
			}
		}
	}
	if found {
		return mi
	}
	return 0
}
