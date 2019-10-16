package main

import "z01"

func main() {

	for y := rune('z'); y >= rune('a'); y-- {

		z01.PrintRune(y)

	}

	z01.PrintRune('\n')

}
