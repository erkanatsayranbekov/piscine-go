package piscine

import "github.com/01-edu/z01"

func Sort(table []int) []int {
	l := 0
	for range table {
		l++
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n > 0 {
		var array []int
		for i := n; i > 0; i = i / 10 {
			array = append(array, i%10)
		}
		array = Sort(array)
		for _, i := range array {
			z01.PrintRune(rune(i + 48))
		}
	} else if n == 0 {
		z01.PrintRune('0')
	}
}
