package main

import (
	"os"

	"github.com/01-edu/z01"
)

func strTOint(s string) int {
	num := 0
	for _, r := range s {
		if '0' <= r && r <= '9' {
			start := 0
			for i := '1'; i <= r; i++ {
				start++
			}
			num = num*10 + start
		}
	}
	return num
}

func main() {
	program := os.Args
	len := len(program)
	Cap := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Low := []rune("abcdefghijklmnopqrstuvwxyz")

	result := []rune{}
	if len == 1 {
		z01.PrintRune(' ')
	} else {
	if program[1] == "--upper" {
		for i := 2; i < len; i++ {
			num := strTOint(program[i])
			if num > 26 || num < 1 {
				result = append(result, rune(32))
			} else {
				result = append(result, Cap[num-1])
			}
		}
	} else {
		for i := 1; i < len; i++ {
			num := strTOint(program[i])
			if num > 26 || num < 1 {
				result = append(result, rune(32))
			} else {
				result = append(result, Low[num-1])
			}
		}
	}
}
	for _, j := range result {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
