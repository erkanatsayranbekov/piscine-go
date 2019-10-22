package piscine

func AlphaCount(str string) int {
	result := 0
	for _, element := range str {
		if element >= 'A' && element <= 'Z' || element >= 'a' && element <= 'z' {
			result++
		}
	}
	return result
}
