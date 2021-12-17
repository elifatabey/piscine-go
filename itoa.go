package piscine

func Itoa(n int) string {
	var final []byte
	var r []byte
	var next int
	var right int

	for {
		if n < 0 {
			n = -1 * n
			r = append(r, '-')
		}
		next = n / 10
		right = n - next*10
		n = next
		final = append(final, byte('0'+right))
		if n == 0 {
			break
		}
	}
	for j := 0; j < len(final); j++ {
		r = append(r, final[len(final)-j-1])
	}
	return string(r)
}
