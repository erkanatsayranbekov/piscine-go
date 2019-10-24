package piscine

func IsUpper(str string) bool {
	var result bool
	for _, j := range str {
		if j >= 'A' && j <= 'Z' {
			result = true
		} else {
			return false
		}
	}
	return result
}
