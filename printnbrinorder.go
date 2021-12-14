package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nums []int
	if n < 0 {
		return
	} else if n == 0 {
		nums = append(nums, 0)
	}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		for _, digits := range nums {
			z01.PrintRune(rune(digits + 48))
		}
	}
}
