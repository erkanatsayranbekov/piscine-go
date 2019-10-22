package piscine

func NRune(s string, n int) rune {
	l := 0
	for range s {
		l++
	}
	if n <= l && n > 0 && l > 0 {
		rune := []rune(s)
		return rune[n-1]
	}
	return '\x00'
}
