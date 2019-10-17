package piscine

func Swap(a *int, b *int) {
	p := *a
	d := *b
	*a = d
	*b = p

}
