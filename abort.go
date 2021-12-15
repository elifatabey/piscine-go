package piscine

func Abort(a, b, c, d, e int) int {
	var collect [5]int
	collect[0] = a
	collect[1] = b
	collect[2] = c
	collect[3] = d
	collect[4] = e

	for i := 0; i < 4; i++ {
		if collect[i] > collect[i+1] {
			collect[i], collect[i+1] = collect[i+1], collect[i]
		}
	}
	return collect[2]
}
