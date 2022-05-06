package main

import "fmt"

func main() {
	name := "Sahal M"

	switch name {
	case "Sahal":
		fmt.Println("Hello Sahal")
	case "Mahfudh":
		fmt.Println("Hello Mahfudh")
	default:
		fmt.Println("Hi,", name)
	}

	// switch bisa tanpa kondisi atau vairable kondisi bisa langsung di letakkan pada case
}
