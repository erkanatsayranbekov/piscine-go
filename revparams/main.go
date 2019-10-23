package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	l := 0
	for range os.Args[1:] {
		l++
	}
	s := os.Args[0]
	for i := l; i > 0; i-- {
		s = os.Args[i]
		for _, j := range s {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
