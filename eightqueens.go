package piscine

import (
	"github.com/01-edu/z01"
)

var result [8]int

func EightQueens() {
	for row := 1; row <= 8; row++ {
		for col := 1; col <= 8; col++ {
			if col == result[row] {
				z01.PrintRune(rune(col))
			}
		}
		z01.PrintRune('\n')
	}
}
func queensave(row int, col int) {
	result[row] = col

}
