package piscine

func LowerRune(s rune) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}
	return false
}

func IsLower(str string) bool {
	for _, value := range str {
		if !LowerRune(value) {
			return false
		}
	}
	return true
}
