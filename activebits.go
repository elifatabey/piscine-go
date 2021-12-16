package piscine

func ActiveBits(n int) int {
	base := []int{128, 64, 32, 16, 8, 4, 2, 1}
	sequence := []int{}

	count := 0
	for _, nums := range base {
		sequence = append(sequence, n/nums) // n/nums if nums is higher it will append 0
		n = n % nums                        // it only divides if result is higher than 0
		// fmt.Println(n)                      // otherwise leaves it as it is (n = kalana esit)
	}
	for _, each := range sequence {
		if each == 1 {
			count++
		}
	}
	return count
}

// func ActiveBits(n int) int {
// 	index := 0
// 	for n > 0 {
// 		if n%2 != 0 {
// 			index++
// 		}
// 		n = n / 2
// 	}
// 	var c int = int(uint(index))
// 	return c
// }
