package main

import (
	"fmt"
)

func StrLen(str string) {
	a := len(str)
	fmt.Println(a)
}

func main() {
	str := "Hello World!"
	StrLen(str)
}
