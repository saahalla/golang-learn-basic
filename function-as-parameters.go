package main

import "fmt"

// use type declaration

type Filter func(string) string

func helloWithFilter(name string, filter Filter) string {
	filteredName := filter(name)
	return "Hello " + filteredName
}

func filterData(name string) string {
	if name == "Ayam" {
		return "..."
	} else {
		return name
	}
}

func main() {
	fmt.Println(helloWithFilter("Sahal", filterData))
	fmt.Println(helloWithFilter("Ayam", filterData))
}
