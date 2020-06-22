package leetcode_go

func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}

	l, h := 0, len(height)-1

	mx := 0

	for l < h {
		mx = max(mx, min(height[l], height[h])*(h-l))
		if height[l] <= height[h] {
			l++
		} else {
			h--
		}
	}

	return mx
}

//func min(x ,y int) int {
//	if x < y{
//		return x
//	}
//	return y
//}
//
//func max(x ,y int) int {
//	if x < y{
//		return y
//	}
//	return x
//}
