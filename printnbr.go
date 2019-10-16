package piscine

import (
	"github.com/01-edu/z01"
)

func Help(n int) {
	if n == 0 {
		return
	} else {
		charX := '0'
		for x := 0; x > n%10; x-- {
			charX++
		}
		Help(n / 10)
		z01.PrintRune(charX)
	}
}

func HelpForPositive(n int) {
	if n == 0 {
		return
	} else {
		charX := '0'
		for x := 0; x < n%10; x++ {
			charX++
		}
		HelpForPositive(n / 10)
		z01.PrintRune(charX)

	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune(45)
		Help(n)
	} else if n == 0 {
		z01.PrintRune(48)
	} else {
		HelpForPositive(n)
	}
}
