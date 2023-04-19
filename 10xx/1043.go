package _0xx

import "fmt"

// 给你一个整数数组 arr，请你将该数组分隔为长度 最多 为 k 的一些（连续）子数组。分隔完成后，每个子数组的中的所有值都会变为该子数组中的最大值。
// 返回将数组分隔变换后能够得到的元素最大和。本题所用到的测试用例会确保答案是一个 32 位整数。
func maxSumAfterPartitioning(arr []int, k int) int {
	sum := make([]int, len(arr))
	maxa := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		maxa[i] = make([]int, len(arr))
		maxa[i][i] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			maxa[i][j] = max(arr[j], maxa[i][j-1])
		}
	}

	for i, v := range arr {
		sm := v
		for j := i - k + 1; j <= i; j++ {
			if j >= 0 {
				var ma = maxa[j][i]
				var tmp = 0
				if j > 0 {
					tmp = sum[j-1]
				}
				if r := ma*(i-j+1) + tmp; r > sm {
					sm = r
				}
			}
		}
		sum[i] = sm
	}

	fmt.Println(sum)

	return sum[len(sum)-1]
}

func max(arr ...int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}
