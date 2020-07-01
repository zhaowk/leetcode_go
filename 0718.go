package leetcode_go

func findLength(A []int, B []int) int {

	mx := 0

	for i := 0; i < len(A); i++ {
		l := 0
		for l < len(B) && len(B)-l > mx {
			j := l
			for ; j < len(B); j++ {
				if B[j] == A[i] {
					break
				}
			}

			if j < len(B) { // found
				k := 1
				for i+k < len(A) && j+k < len(B) && A[i+k] == B[j+k] {
					k++
				}
				mx = max(mx, k)
				l = j + 1
			} else {
				break
			}
		}
	}
	return mx
}
