package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	sep := ' '
	last := 0
	for i, e := range str {
		if e == sep {
			result = append(result, string(str[last:i]))
			last = i + 1
		}

	}
	result = append(result, string(str[last:]))
	return result
}
