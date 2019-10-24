package piscine

func NRune(s string, n int) rune {
	rune := []rune(s)
	return rune[n-1]
}
