package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := len(a) - 1; i >= 1; i-- {
		for _, arg := range a[i] {
			z01.PrintRune(arg)
		}
		z01.PrintRune('\n')
	}
}
