package piscine

func SortIntegerTable(table []int) {
	a := 0
	for i := range table {
		a = i
		a++
	}

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
