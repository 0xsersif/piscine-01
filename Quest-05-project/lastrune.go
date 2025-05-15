package piscine

func LastRune(s string) rune {
	x := []rune(s)
	if len(x) == 0 {
		return 0
	}
	return x[len(x)-1]
}
