package _9xx

func distinctSubseqII(s string) int {
	const mod = 1e9 + 7
	var repeated [26]int
	cnt := 1
	for _, c := range s {
		tmp := cnt
		cnt = (2*cnt - repeated[c-'a'] + mod) % mod
		repeated[c-'a'] = tmp
	}
	return (cnt - 1) % mod
}
