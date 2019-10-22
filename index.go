package piscine

func Len(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func Index(s string, toFind string) int {
	result := -1
	StrLen := Len(s)
	FinLen := Len(toFind)
	arrS := []rune(s)
	arrF := []rune(toFind)
	if s == "" {
		return 0
	}
	for index, letter := range s {
		if letter == arrF[0] && index+FinLen <= StrLen {
			j := 0
			c := 0
			for i := index; i < FinLen; i++ {
				if arrS[i] == arrF[j] {
					c++
				}
				j++
			}
			if c == j {
				result = index
				break
			}
		}
	}
	return result
}
