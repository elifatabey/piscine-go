package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		s := strconv.FormatBool(isPowerOf2(nbr))
		for _, char := range s {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
