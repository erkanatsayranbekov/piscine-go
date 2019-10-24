package piscine

func IsAlpha(str string) bool {
	var result bool
	for _, j := range str {
		if j >= 'a' && j <= 'z' || j >= 'A' && j <= 'Z' || j >= '0' && j <= '9' {
			result = true
		} else {
			return false
		}
	}
	return result
}
