package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	counter := 0
	for range array {
		counter++
	}
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if f(array[i], array[j]) > 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
