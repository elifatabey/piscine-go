package main

import (
	"github.com/01-edu/z01"
)

func main() {
	num := []rune("0123456789")
	i := 0
	for i <= 9 {
		z01.PrintRune(num[i])
		i++
	}

	z01.PrintRune('\n')
}
