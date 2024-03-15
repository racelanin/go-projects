package piscine

func StringToIntSlice(str string) []int {
	var arr []int

	for _, r := range str {
		arr = append(arr, int(r))
	}

	return arr
}

//func main() {
//	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
//	fmt.Println(StringToIntSlice("Converted this string into an int"))
//	fmt.Println(StringToIntSlice("hello THERE"))
//}
