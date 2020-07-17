package leetcode_go

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

func intersect2(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int, 0)

	for i := range nums1 {
		m[nums1[i]]++
	}

	for i := range nums2 {
		if m[nums2[i]] > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}
