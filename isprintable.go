package piscine

func IsPrintable(str string) bool {
	var result bool
	for _, j := range str {
		if j > 31 {
			result = true
		} else {
			return false
		}
	}
	return result
}
