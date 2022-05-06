package main

import "fmt"

func main() {
	var books map[string]string = map[string]string{
		"title": "Belajar GoLang",
		"year":  "2020",
		"city":  "Sleman",
	}
	books["authors"] = "Sahal"
	fmt.Println(books)

	persons := make(map[string]string)
	persons["name"] = "Sahal"
	persons["address"] = "Sleman"
	persons["job"] = "Software Engineer"

	fmt.Println(persons)
}
