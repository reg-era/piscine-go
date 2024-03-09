package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	res := []int{}

	for i := 0; i < len(slice1)/2; i++ {
		slice1[i], slice1[len(slice1)-1-i] = slice1[len(slice1)-1-i], slice1[i]
	}

	for i := 0; i < len(slice2)/2; i++ {
		slice2[i], slice2[len(slice2)-1-i] = slice2[len(slice2)-1-i], slice2[i]
	}

	if len(slice1) == 0 {
		return slice2
	} else if len(slice2) == 0 {
		return slice1
	}

	if len(slice1) == len(slice2) {
		for i := 0; i < len(slice1)*2; i++ {
			if i%2 == 0 {
				res = append(res, slice1[i/2])
			} else {
				res = append(res, slice2[i/2])
			}
		}
		return res
	}

	if len(slice1) > len(slice2) {
		dim := len(slice1)
		ind := 0
		for dim != len(slice2) {
			res = append(res, slice1[ind])
			ind++
			dim--
		}

		nwslc := []int{}
		for i := ind; i < len(slice1); i++ {
			nwslc = append(nwslc, slice1[i])
		}

		for i := 0; i < len(slice2)*2; i++ {
			if i%2 == 0 {
				res = append(res, nwslc[i/2])
			} else {
				res = append(res, slice2[i/2])
			}
		}

		return res
	}

	if len(slice2) > len(slice1) {
		dim := len(slice2)
		ind := 0
		for dim != len(slice1) {
			res = append(res, slice2[ind])
			ind++
			dim--
		}

		nwslc := []int{}
		for i := ind; i < len(slice2); i++ {
			nwslc = append(nwslc, slice2[i])
		}

		for i := 0; i < len(slice1)*2; i++ {
			if i%2 == 0 {
				res = append(res, slice1[i/2])
			} else {
				res = append(res, nwslc[i/2])
			}
		}

		return res
	}

	return res
}
