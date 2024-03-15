package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	result := ""
	word := ""
	count := 0
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			word += string(str[i])
			count++
			if count == 5 {
				result += word + " "
				word = ""
				count = 0
				i += 2
				continue
			}
		}
		i++
	}
	if word != "" {
		result += word
	}
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
