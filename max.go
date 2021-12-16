package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		var newa []int = a
		for i := 0; i < len(a)-1; i++ {
			for j := 0; j < len(a)-1; j++ {
				if newa[j] < newa[j+1] {
					newa[j], newa[j+1] = newa[j+1], newa[j]
				}
			}
		}
		return newa[0]
	}
}
