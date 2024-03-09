package piscine

func ConcatParams(args []string) string {
	res := make([]string, len(args))
	str := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			res[i] = args[i]
			break
		}
		res[i] = args[i] + "\n"
	}

	for i := 0; i < len(res); i++ {
		str += res[i]
	}

	return str
}
