package leetcode_go

func convertToTitle(n int) string {
	var b []byte

	for n > 0 {
		b = append(b, byte((n-1)%26+'A'))
		n = (n - 1) / 26
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return string(b)
}
