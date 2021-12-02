package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	num := []rune("0123456789")
	for i := 0; i <= 7; i++ {
		for j := i + 1; i <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				if i == 7 && j == 8 && k == 9 {
					z01.PrintRune(num[i])
					z01.PrintRune(num[j])
					z01.PrintRune(num[k])
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(num[i])
					z01.PrintRune(num[j])
					z01.PrintRune(num[k])
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
