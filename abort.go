package piscine

func Abort(a, b, c, d, e int) int {
	var collect [5]int
	collect[0] = a
	collect[1] = b
	collect[2] = c
	collect[3] = d
	collect[4] = e

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if collect[i] > collect[j] {
				collect[i], collect[j] = collect[j], collect[i]
			}
		}
	}
	return collect[2]
}
