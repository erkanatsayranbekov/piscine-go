package piscine

func Number(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func IsNumeric(str string) bool {
	for _, value := range str {
		if !Number(value) {
			return false
		}
	}
	return true
}
