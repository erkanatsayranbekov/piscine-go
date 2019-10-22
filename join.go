package piscine

func Join(ss []string, sep string) string {
	str := ""
	for _, value := range ss {
		if str != "" {
			str = str + sep
		}
		str = str + value
	}
	return str
}
