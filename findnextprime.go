package piscine

func FindNextPrime(nb int) int {
	if nb == 1 || nb <= 0 {
		return 2
	} else {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				return FindNextPrime(nb + 1)
			}
		}
	}
	return nb
}
