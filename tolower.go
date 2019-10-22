package piscine

func ToLower(s string) string {
	low := []rune(s)
	for i, letter := range low {
		if letter >= 'A' && letter <= 'Z' {
			low[i] = letter + 32
		}
	}
	result := string(low)
	return result
}
