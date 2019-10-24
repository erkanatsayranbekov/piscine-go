package piscine

func BasicJoin(strs []string) string {
	str := ""
	for _, i := range strs {
		str = str + i
	}
	return str
}
