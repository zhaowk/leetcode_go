package leetcode_go

func countDigitOne(n int) (ans int) {
	for k := 1; n >= k; k *= 10 {
		ans += (n/(k*10))*k + min(max(n%(k*10)-k+1, 0), k)
	}
	return
}
