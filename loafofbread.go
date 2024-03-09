package piscine

func LoafOfBread(str string) string {

	if len(str) == 0 {

		return "\n"

	} else if len(str) < 5 {

		return "Invalid Output\n"
	} else {
		result := ""
		spacesCount := 0
		for _, char := range str {

			if spacesCount%6 != 5 && char == ' ' {
				continue
			}
			if spacesCount%6 == 5 {
				result += " "
			} else {
				result += string(char)
			}
			spacesCount++
		}
		for i := len(result) - 1; i >= 10; i-- {
			if result[i] != ' ' {
				result = result[:1+1]
				break
			}
		}
		return result + "\n"
	}
}
