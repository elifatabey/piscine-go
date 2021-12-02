package piscine

func Swap(a *int, b *int) {
	var cb int = *b
	var ca int = *a
	*a = cb
	*b = ca
}
