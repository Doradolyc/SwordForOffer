package main

// 时间复杂度 o(n * m)
// 空间复杂度 o(1)
//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	for i:= 0; i < len(matrix); i++ {
//		for j:= 0; j < len(matrix[i]); j++ {
//			if matrix[i][j] == target {
//				return true
//			}
//		}
//	}
//	return false
//}

// 时间复杂度 o(n + m)
// 空间复杂度 o(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	columns := len(matrix[0])

	r := 0
	c := columns - 1

	for r < row && c >= 0 {
		num := matrix[r][c]

		if num == target {
			return true
		} else if num > target {
			c--
		} else if num < target {
			r++
		}
	}

	return false
}