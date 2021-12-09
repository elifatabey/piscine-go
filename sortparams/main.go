package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	program := os.Args

	for i := len(program); i > 0; i-- {
		for j := 2; j < i; j++ {
			if program[j-1] > program[j] {
				middle := program[j]
				program[j] = program[j-1]
				program[j-1] = middle
			}
		}
	}
	for i := 0; i < len(program); i++ {
		for _, r := range program[i] {
			// if i > 1 {
			z01.PrintRune(r)
			// }
		}
		z01.PrintRune('\n')
	}
}
