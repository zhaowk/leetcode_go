package leetcode_go

func canArrange(arr []int, k int) bool {

	la := len(arr)
	if la == 0 {
		return false
	}

	for i := 0; i < la-1; i += 2 {
		found := false
		for j := i + 1; j < la; j++ {
			if (arr[j]+arr[i])%k == 0 {
				found = true
				arr[i+1], arr[j] = arr[j], arr[i+1]
			}
		}
		if !found {
			return false
		}
	}
	return true
}
