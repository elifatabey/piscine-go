package main

import "github.com/01-edu/z01"

func main() {
	alph := []rune("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < 26; i++ {
		z01.PrintRune(alph[i])
	}
	z01.PrintRune('\n')
}
