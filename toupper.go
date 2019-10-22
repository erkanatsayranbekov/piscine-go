package piscine

func ToUpper(s string) string {
	upper := []rune(s)
	for i, letter := range upper {
		if letter >= 'a' && letter <= 'z' {
			upper[i] = letter - 32
		}
	}
	result := string(upper)
	return result
}
