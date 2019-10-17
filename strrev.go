package piscine

func StrLen(str string) int {
	a := 0
	s := []rune(str)
	for i := range s {
		a = i
	}
	return a
}

func StrRev(s string) string {
	j := 0
	nrune := []rune(s)
	for i := StrLen(s); i >= 0; i-- {
		nrune[j] = rune(s[i])
		j++
	}
	return string(nrune)
}
