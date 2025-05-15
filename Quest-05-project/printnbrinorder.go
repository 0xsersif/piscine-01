package piscine

import "github.com/01-edu/z01"

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	digits = SortIntegerTable(digits)

	for _, digit := range digits {
		z01.PrintRune(rune(digit + '0'))
	}
}
