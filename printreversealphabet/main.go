package main

import (
	"github.com/01-edu/z01"
)

func main() {
	alph := []rune("abcdefghijklmnopqrstuvwxyz")
	for i := 25; i >= 0; i-- {
		z01.PrintRune(alph[i])
	}
	z01.PrintRune('\n')
}
