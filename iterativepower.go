package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || nb < 0 {
		return 0
	} else {
		result := nb
		for i := 1; i < power; i++ {
			result = nb * result
		}
		return result
	}
}
