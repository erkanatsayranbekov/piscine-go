package piscine

func LastRune(s string) rune {
	l := 0
	for range s {
		l++
	}
	rune := []rune(s)
	return rune[l-1]
}
