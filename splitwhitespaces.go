package piscine

func SplitWhiteSpaces(str string) []string {
	l := 0
	for _, i := range str {
		if i == ' ' || i == '	' || i == '\n' {
			l++
		}
	}
	result := make([]string, l+1)
	last := 0
	m := 0
	for i, e := range str {
		if e == ' ' && i != last || e == '	' && i != last || e == '\n' && i != last {
			result[m] = string(str[last:i])
			last = i + 1
			m++
		} else {
			continue
		}
	}
	result[l] = string(str[last:])
	return result
}
