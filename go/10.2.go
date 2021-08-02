package main

// 时间 o(n)
// 空间 o(n)
func numWays(n int) int {
	var num []int

	num = append(num, 1)

	// 一层就一种上法
	num = append(num, 1)
	// 两层两种上法
	num = append(num, 2)

	for i := 3; i <= n; i++ {
		num = append(num, (num[i - 1] + num[i - 2]) % 1000000007)
	}

	return num[n]
}
