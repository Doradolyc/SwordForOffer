package main

// 时间复杂度o(n)
// 空间复杂度o(1)
//func replaceSpace(s string) string {
//	return strings.ReplaceAll(s, " ", "%20")
//}


// 时间复杂度o(n)
// 空间复杂度o(n)
func replaceSpace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i + 1:]
			i += 2
		}
	}

	return s
}
