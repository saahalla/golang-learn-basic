package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	newSlice := make([]string, 2, 5)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var iniArray = [4]int{1, 2, 3, 4}
	var iniSlice = []int{1, 2, 3, 4}

	fmt.Println(iniArray, iniSlice)
}
