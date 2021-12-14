package piscine

func IterativeFactorial(nb int) int {
	count := 1
	if nb > 20 || nb < 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			count = count * i
		}
	}
	return count
}
