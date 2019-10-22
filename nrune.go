package main

import (
	"github.com/01-edu/z01"
)

func NRune(s string, n int) rune {
	l := 0
	for range s {
		l++
	}
	if n <= l && n > 0 && l > 0 {
		rune := []rune(s)
		return rune[n-1]
	}
	return '\x00'
}

func main() {
	z01.PrintRune(NRune("Hello!", -3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
