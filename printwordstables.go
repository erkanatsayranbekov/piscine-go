package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, j := range table {
		for _, i := range j {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
