package leetcode_go

//  四平方和定理
//  4 ** k + ( 8 * k + 7)  ==> 四平方和, 4   (n/4/4/4/4...) % 8 == 7
//  x 平方数                ==> 1 (x**2 == n)
//  两平方数之和             ==> 2 (x**2 + y**2 == n)
//  others                 ==> 3



// 垃圾动态规划
func numSquares(n int) int {
	m := make(map[int]int)
	a := make([]int, 101) // max
	r := make([]int, n+1)
	for i := int(1); i <= 100; i++ {  // 0, 1, 4, 9, ..., 10000(max)
		a[i] = i * i
		m[i*i] = i
	}

	r[0] = 0
	r[1] = 1
	var idx = 1

	for i := 2; i <= n ;i++ {
		if m[i] > 0 {
			r[i] = 1  // target
			idx = m[i]
		} else {
			r[i] = 1 + r[i-1]
			for j := 2; j <= idx && i - a[j] >= 0; j++ {
				r[i] = min(r[i], r[i - a[j]] + 1)
			}
		}
	}

	//fmt.Println(r)

	return r[n]
}

