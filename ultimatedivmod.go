package piscine

func UltimateDivMod(a *int, b *int) {
	p := *a
	d := *a / *b
	*b = p % *b
	*a = d
}
