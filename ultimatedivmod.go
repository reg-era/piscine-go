package piscine

func UltimateDivMod(a *int, b *int) {
	x, y := *a, *b

	resdiv, resmod := x/y, x%y

	*a, *b = resdiv, resmod
}
