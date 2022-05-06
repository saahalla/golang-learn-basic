package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

// function with return
func tambah(a int, b int) string {
	c := a + b
	// return "Jumlah dari " + string(a) + string(b) + " = " + string(c) // ini salah :D
	return fmt.Sprintf("Jumlah dari %v + %v = %v", a, b, c)
}

// function return multiple value and named value
func getNames() (firstName string, middleName string, lastName string) {
	firstName = "Sahal"
	middleName = "Mahfudh"
	lastName = "Aput"

	return
}

func main() {
	sayHello("Mahfudh", "Sahal")
	sayHello("Mahfudh", "Sahal")
	sayHello("Mahfudh", "Sahal")

	fmt.Println(tambah(1, 2))
	fn, mn, ln := getNames()
	fmt.Println(getNames())
	fmt.Println(fn)
	fmt.Println(mn)
	fmt.Println(ln)
}
