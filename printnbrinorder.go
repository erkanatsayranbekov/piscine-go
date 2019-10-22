package piscine

import "github.com/01-edu/z01"

func Sort(table []int) []int {
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
	return table
}

func PrintNbrInOrder(n int) {
	if n > 0 {
		array := []int{}
		for i := n; i > 0; i /= 10 {
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
