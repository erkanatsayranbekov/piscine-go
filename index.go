package piscine

func Len(table []rune) int {
	len := 0
	for range table {
		len++
	}
	return len
}

func Index(s string, toFind string) int {
	as := []rune(s)
	af := []rune(toFind)
	ls := Len(as)
	lf := Len(af)
	index := -1
	count := 0
	if lf > 0 {
		for j, i := range as {
			if i == af[0] && j+lf <= ls && count == 0 {
				for k := j; k <= j+lf; k++ {
					for m := 0; m <= lf-1; m++ {
						if as[k] == af[m] {
							index = j
							count++
						}
					}
				}
			}
		}
		return index
	}
	return 0
}
