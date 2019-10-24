package piscine

func IsLower(str string) bool {
	var result bool
	for _, j := range str {
		if j >= 'a' && j <= 'z' {
			result = true
		} else {
			return false
		}
	}
	return result
}
