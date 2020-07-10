package leetcode_go

func divingBoard(shorter int, longer int, k int) []int {

	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{k * shorter}
	}

	result := make([]int, k+1)

	for i := 0; i <= k; i++ {
		result[i] = i*longer + (k-i)*shorter
	}
	return result
}
