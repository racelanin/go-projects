package piscine

func ReverseMenuIndex(menu []string) []string {
	// for i, j := 0, len(menu)-1; i < j; i, j = i+1, j-1 {
	//     menu[i], menu[j] = menu[j], menu[i]
	// }
	menuLen := len(menu)
	result := make([]string, menuLen)

	for i, n := range menu {
		j := menuLen - i - 1

		result[j] = n
	}
	return result
}
