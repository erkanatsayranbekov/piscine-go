package piscine

func UpperRune(s rune) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}

func IsUpper(str string) bool {
	for _, value := range str {
		if !UpperRune(value) {
			return false
		}
	}
	return true
}
