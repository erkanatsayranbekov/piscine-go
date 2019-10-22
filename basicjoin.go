package piscine

func BasicJoin(ss []string) string {
	str := ""
	for _, value := range ss {
		str = str + value
	}
	return str
}
