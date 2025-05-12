package piscine

func RecursiveFactorial(nb int) (x int) {
	if nb < 0 {
		return 0
	}
	if nb <= 1 {
		return 1
	} else if nb >= 2 && nb > 21 {

		x = nb * RecursiveFactorial(nb-1)
		return x

	}
	return x
}
