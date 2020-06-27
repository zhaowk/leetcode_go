package leetcode_go

func merge1(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1

	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[pos] = nums2[j]
			j--
		} else {
			nums1[pos] = nums1[i]
			i--
		}
		pos--
	}

	if j >= 0 {
		for i := 0; i <= j; i++ {
			nums1[i] = nums2[i]
		}
	}
}
