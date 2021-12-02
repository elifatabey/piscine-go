package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	num := []rune("0123456789")
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for m := 0; m <= 9; m++ {
					if i*10+j >= k*10+m {
					} else {
						if i == 9 && j == 8 && k == 9 && m == 9 {
							z01.PrintRune(num[i])
							z01.PrintRune(num[j])
							z01.PrintRune(' ')
							z01.PrintRune(num[k])
							z01.PrintRune(num[m])
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(num[i])
							z01.PrintRune(num[j])
							z01.PrintRune(' ')
							z01.PrintRune(num[k])
							z01.PrintRune(num[m])
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
