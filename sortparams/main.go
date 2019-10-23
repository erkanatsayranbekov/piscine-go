package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Len(table []string) int {
	len := 0
	for range table {
		len++
	}
	return len
}

func SortTable(table []string) []string {
	a := Len(table)
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
	s := os.Args[1:]
	SortTable(s)
	l := Len(s)
	for a := 0; a < l; a++ {
		q := s[a]
		for _, j := range q {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}
