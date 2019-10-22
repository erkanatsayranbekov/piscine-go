package piscine

func LetterOrNumber(s rune) bool {
	if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' || s >= '0' && s <= '9' {
		return true
	}
	return false
}

func IsAlpha(str string) bool {
	for _, value := range str {
		if !LetterOrNumber(value) {
			return false
		}
	}
	return true
}
