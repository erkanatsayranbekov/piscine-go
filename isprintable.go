package piscine

func IsPrintable(str string) bool {
	for _, value := range str {
		if value < 32 {
			return false
		}
	}
	return true
}
