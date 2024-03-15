package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	item := make(map[string]int)

	for _, v := range Space(str) {
		item[v]++
	}

	return item
}

func Space(s string) []string {
	space := make([]string, 0, len(s))
	word := ""

	for _, v := range s {
		if v == ' ' {
			space = append(space, word)
			word = ""
		} else {
			word += string(v)
		}
	}
	space = append(space, word)

	return space
}

//func main() {
//	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
//	for index, element := range ShoppingSummaryCounter(summary) {
//		fmt.Println(index, "=>", element)
//	}
//}
