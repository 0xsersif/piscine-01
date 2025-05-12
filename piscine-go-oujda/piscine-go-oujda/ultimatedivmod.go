package piscine

func UltimateDivMod(a *int, b *int) {
	div := 0
	div = *a
	*a = div / *b
	*b = div % *b
}
