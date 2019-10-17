package piscine

func UltimateDivMod(a *int, b *int) {
	p := *a
	*a = *a / *b
	*b = p % *b
}
