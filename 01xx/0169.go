package _1xx

func majorityElement(nums []int) int {
	target, cnt := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if target != nums[i] {
			cnt--
			if cnt == 0 {
				target, cnt = nums[i], 1
			}
		} else {
			cnt++
		}
	}
	return target
}
