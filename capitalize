package piscine

func Number(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func LowerRune(s rune) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}
	return false
}

func UpperRune(s rune) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}

func LetterOrNumber(s rune) bool {
	if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' || s >= '0' && s <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	arr := []rune(s)
	str := []rune(s)
	if LowerRune(str[0]) {
		arr[0] = str[0] - 32
	}
	for i := range str {
		if i != 0 && LowerRune(str[i]) && !LetterOrNumber(str[i-1]) {
			arr[i] = str[i] - 32
		} else if i != 0 && UpperRune(str[i]) && (UpperRune(arr[i-1]) || LowerRune(arr[i-1]) || Number(arr[i-1])) {
			arr[i] = str[i] + 32
		}
	}
	result := string(arr)
	return result
}
