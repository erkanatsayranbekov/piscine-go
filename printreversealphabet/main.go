package main

import (
	"github.com/01-edu/z01"
)

func main() {

	for y := 'z'; y >= 'a'; y-- {
		z01.PrintRune(y)
		if y == 'a' {
			z01.PrintRune(10)
		}
	}
}
