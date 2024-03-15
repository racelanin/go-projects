package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	middle := len(arr) / 2
	if len(arr)%2 == 1 {
		return arr[middle]
	} else {
		return (arr[middle-1] + arr[middle]) / 2
	}
}

//func main() {
//	middle := Abort(2, 3, 8, 5, 7)
//	fmt.Println(middle)
//}
