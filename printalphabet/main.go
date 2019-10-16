package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := rune('a'); i <= rune('z'); i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
