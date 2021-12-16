package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		var newa []int = a
		for i := 0; i < len(a)-1; i++ {
			if newa[i] < newa[i+1] {
				newa[i], newa[i+1] = newa[i+1], newa[i]
			}
		}
		return newa[0]
	}
}
