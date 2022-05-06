package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

/**
contoh function menggunakan param struct
*/
func sayHi(customer Customer, name string) {
	fmt.Println("Hai", name, "My name is", customer.Name)
}

/**
contoh struct function
*/
func (customer Customer) sayHai(name string) {
	fmt.Println("Hai", name, "My name is", customer.Name)
}

func main() {
	sahal := Customer{}
	sahal.Name = "Sahal Mahfudh"
	sahal.Address = "Sleman, DIY"
	sahal.Age = 20

	// fmt.Println(sahal)
	// struct function / struct methode
	sayHi(sahal, "aput") // function biasa
	sahal.sayHai("Mahfudh")
}
