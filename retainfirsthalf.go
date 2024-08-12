package piscine

func RetainFirstHalf(str string) string {
	if len(str) == 1 {
		return str
	} else if len(str)%2 == 0 {
		return str[:len(str)/2]
	} else if len(str)%2 != 0 {
		return str[:len(str)/2]
	} else {
		return ""
	}
}
