package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Counter", i)
	}

	slice := []string{"Sahal", "Mahfudh", "Aput"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("============ For Range =============")

	for i, value := range slice {
		fmt.Println("Index", i, "Value", value)
	}
}
