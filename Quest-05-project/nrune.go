package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	if n > len(x) || n < 1 {
		return 0
	}
	return x[n-1]
}
