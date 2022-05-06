package main

import "fmt"

func sumNumber(numbers ...int) (sum int) {
	sum = 0

	for _, val := range numbers {
		sum += val
	}

	return
}

func main() {
	total := sumNumber(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	// with slice
	totalSlice := sumNumber([]int{1, 2, 3, 4, 5}...)
	fmt.Println(totalSlice)
}
