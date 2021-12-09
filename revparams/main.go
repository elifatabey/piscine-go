package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	program := os.Args

	for i := len(program) - 1; i >= 1; i-- {
		for _, r := range program[i] {
			// if i > 1 {
			z01.PrintRune(r)
			// }
		}
		z01.PrintRune('\n')
	}
}
