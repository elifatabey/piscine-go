package piscine

func UltimateDivMod(a *int, b *int) {
	var c int = *a / *b
	*b = *a % *b
	*a = c
}
