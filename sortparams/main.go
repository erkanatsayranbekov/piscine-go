package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable(table []string) []string {
	a := 0
	for range table {
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
func main() {
	len := 0
	s := os.Args[1:]
	SortIntegerTable(s)
	for range s {
		len++
	}
	for a := 1; a <= len; a++ {
		q := s[a]
		for _, j := range q {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}
