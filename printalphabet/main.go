package main

import (
	"github.com/01-edu/z01"
)

func main() {
	alph := []rune("abcdefjhijklmnopqrstuvwxyz")
	for i := 0; i < 26; i++ {

		z01.PrintRune(alph[i])

	}
	z01.PrintRune('\n')
}
