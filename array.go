package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Sahal"
	names[1] = "Mahfudh"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{1, 2, 3}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
