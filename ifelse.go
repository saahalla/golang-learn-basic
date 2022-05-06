package main

import "fmt"

func main() {
	name := "Sahal M"

	if name == "Sahal" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi, What is your name ?")
	}

	// short statements if

	if length := len(name); length > 5 {
		fmt.Println("Max char = 5")
	} else {
		fmt.Println("Name is OK")
	}
}
