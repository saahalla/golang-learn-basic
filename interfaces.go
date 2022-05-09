package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	/**
	Mohon pelajari lagi ya :)
	yang saya paham, interface bisa menerapkan methode di dalam struct tapi tidak langsung seperti class, bikin struct function baru terlebih dahulu
	- type data abstract
	*/
	var sahal Person
	sahal.Name = "sahal"
	SayHello(sahal)

}
