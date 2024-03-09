package piscine

func DivMod(a int, b int, div *int, mod *int) {
	resdiv := a / b
	resmod := a % b

	*div = resdiv
	*mod = resmod
}
