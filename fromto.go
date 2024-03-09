package piscine

// 1 -> 10
// 10 -> 1

// 01, 02, 03, 04, 05, 06, 07, 08, 09, 10

func FromTo(a, z int) string {
	inv := "Invalid\n"

	if (a > 100) && (a < -1) || (z > 100) && (z < -1) {
		return inv
	}

	if a == z {
		return "\n"
	}

	if a > z {
	}

	if a < z {
	}

	return inv
}
