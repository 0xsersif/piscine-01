package piscine

func UltimateDivMod(a *int, b *int) {
	var div int
	div = *a
	*a = div / *b
	*b = div % *b
}
