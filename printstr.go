package piscine

import (
	"github.com/01-edu/z01"
)

func printstr(s string) {
	convertrune := []rune(s)
	for i := 0; i <= len(convertrune)-1; i++ {
		z01.PrintRune(convertrune[i])
	}
	z01.PrintRune('\n')
}
