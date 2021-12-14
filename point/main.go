package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func intToRune(num int) {
	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	l := len(nums) - 1
	for i := l; i >= 0; i-- {
		z01.PrintRune(rune(nums[i] + 48))
	}
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	intToRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	intToRune(points.y)
	z01.PrintRune('\n')
}
