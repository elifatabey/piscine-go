package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		if power == 1 {
			return nb
		}
		if power > 1 {
			return nb * RecursiveFactorial(nb-1)
		}
		return 1
	}
}
