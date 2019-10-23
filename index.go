package piscine

// import "fmt"

func Index(s string, toFind string) int {
	var rune rune
	for _, i := range toFind {
		rune = i
		break
	}

	index := 0
	for j, i := range s {
		if i == rune {
			index = j
			break
		} else {
			index = -1
		}
	}
	return index
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hO"))
// }
