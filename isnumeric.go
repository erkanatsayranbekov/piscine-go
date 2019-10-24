package piscine

func IsNumeric(str string) bool {
	var result bool
	for _, j := range str {
		if j >= '0' && j <= '9' {
			result = true
		} else {
			return false
		}
	}
	return result
}
