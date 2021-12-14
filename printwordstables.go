package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, string := range a {
		for _, rune := range []rune(string) {
			z01.PrintRune(rune)
		}
		z01.PrintRune('\n')
	}
}
