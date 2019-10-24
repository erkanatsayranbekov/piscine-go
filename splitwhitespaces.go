package piscine

func SplitWhiteSpaces(str string) []string {
	l := 0
	for _, i := range str {
		if i == ' ' && i-1 != ' ' || i == '	' && i-1 != '	' || i == '\n' && i-1 != '\n' {
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
		} else if e == ' ' && i == last || e == '	' && i == last || e == '\n' && i == last {
			last = i + 1
		}
	}
	result[l] = string(str[last:])
	return result
}
