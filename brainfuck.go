package piscine

var (
	Array [2048]byte
	Point int
)

func OpperationBF(code byte) {
	if code == '>' {
		Point++
	}
	if code == '<' {
		Point--
	}
	if code == '+' {
		Array[Point]++
	}
	if code == '-' {
		Array[Point]--
	}
}

func FindEnd(ind int, code string) int {
	res := 0
	for i := len(code) - 1; i >= ind; i-- {
		if code[i] == ']' {
			res = i
			break
		}
	}
	return res
}

func BrainFuck(code string) string {
	res := ""
	for i := 0; i < len(code); i++ {
		if code[i] == '>' || code[i] == '<' || code[i] == '+' || code[i] == '-' || code[i] == '.' || code[i] == '[' || code[i] == ']' {
			OpperationBF(code[i])
			if code[i] == '.' {
				res += string(byte(Array[Point]))
			}
			if code[i] == '[' {
				rang := Array[Point]
				start := i + 1
				end := FindEnd(i, code)
				for rang > 0 {
					BrainFuck(code[start:end])
					rang--
				}
				i = end
			}
		} else {
			continue
		}
	}
	return res
}
