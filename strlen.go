package piscine

func StrLen(str string) int {
	a := 0
	s := []rune(str)
	for i := range s {
		a = i
	}
	return a + 1

}
