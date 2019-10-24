package piscine

func ConcatParams(args []string) string {
	l := 0
	for range args {
		l++
	}
	str := ""
	for j, i := range args {
		if j != l-1 {
			str = str + i
			str += "\n"
		} else {
			str = str + i
		}
	}
	return str
}
