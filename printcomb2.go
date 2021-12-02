package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	num := []rune("0123456789")
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if i == 9 && j == 9 {
				z01.PrintRune(num[i])
				z01.PrintRune(num[j])
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(num[i])
				z01.PrintRune(num[j])
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
