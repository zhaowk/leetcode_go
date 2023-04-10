package _6xx

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sums := make([]int, len(nums)-k+1)

	for i := 0; i < k; i++ {
		sums[0] += nums[i]
	}
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] - nums[i-1] + nums[i+k-1]
	}

	x, y, z, maxi := 0, k, 2*k, sums[0]+sums[k]+sums[2*k]
	sumX, sumXIdx := sums[x], x
	sumXY, sumXYIdxX, sumXYIdxY := sums[x]+sums[y], x, y

	for i := 2 * k; i < len(sums); i++ {
		if sums[i-2*k] > sumX {
			sumX, sumXIdx = sums[i-2*k], i-2*k
		}

		if sumX+sums[i-k] > sumXY {
			sumXY, sumXYIdxX, sumXYIdxY = sumX+sums[i-k], sumXIdx, i-k
		}

		if sumXY+sums[i] > maxi {
			x, y, z, maxi = sumXYIdxX, sumXYIdxY, i, sumXY+sums[i]
		}
	}

	return []int{x, y, z}
}

//func maxSumOfThreeSubarrays(nums []int, k int) []int {
//	x, y, z := 0, k, 2 * k
//	sumX, sumY, sumZ := 0, 0, 0
//	maxi := 0
//
//	for a := 0; a < k; a++ {
//		sumX += nums[a]
//	}
//	for i := 0; i <= len(nums) - 3 * k; i++ {
//		sumY = 0
//		for a := i + k; a < i + 2 * k; a++ {
//			sumY += nums[a]
//		}
//		for j := i + k; j <= len(nums) - 2 * k; j++ {
//			sumZ = 0
//			for a := j + k; a < j + 2 * k; a++ {
//				sumZ += nums[a]
//			}
//			for m := j + k; m <= len(nums) - k; m++ {
//				if sumX + sumY + sumZ > maxi {
//					x, y, z, maxi = i, j, m, sumX + sumY + sumZ
//				}
//				if m + k < len(nums) {
//					sumZ += nums[m+k] - nums[m]
//				}
//			}
//			sumY += nums[j+k] - nums[j]
//		}
//		sumX += nums[i+k] - nums[i]
//	}
//
//	return []int{x, y, z}
//}
