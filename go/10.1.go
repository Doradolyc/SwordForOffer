package main


// 时间 o(n)
// 空间 o(n)
func fib(n int) int {
	var num []int

	num = append(num, 0)
	num = append(num, 1)

	for i:= 2; i <= n; i++ {
		num = append(num, (num[i - 1] + num[i - 2]) % 1000000007)
	}

	return num[n]
}