package piscine

func IsPrintable(str string) bool {
	var result bool
	for _, j := range str {
		if j >= 'a' && j <= 'z' || j >= 'A' && j <= 'Z' {
			result = true
		} else {
			return false
		}
	}
	return result
}
