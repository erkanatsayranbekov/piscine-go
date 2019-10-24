package piscine

func AlphaCount(str string) int {
	result := 0
	for _, i := range str {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			result++
		}
	}
	return result
}
