package piscine

func IsWord(j rune) bool {
	var result bool
	if j >= 'a' && j <= 'z' || j >= 'A' && j <= 'Z' || j >= '0' && j <= '9' {
		result = true
	} else {
		return false
	}
	return result
}

func Capitalize(s string) string {
	r := []rune(s)
	for j, i := range r {
		if j == 0 && i >= 'a' && i <= 'z' || !IsWord(r[j-1]) && i >= 'a' && i <= 'z' {
			r[j] = i - 32
		} else if i >= 'A' && i <= 'Z' && IsWord(r[j-1]) {
			r[j] = i + 32
		}

	}
	return string(r)
}
