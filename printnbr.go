package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		Help(n * -1)
	} else {
		Help(n)
	}
}

func Help(n int) {
	div := n % 10
	var x rune = '0'
	for i := 1; i <= div; i++ {
		x++
	}
	if div == 0 {
		return
	}
	Help(n / 10)
	z01.PrintRune(x)
}
