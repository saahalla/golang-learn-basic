package main

import "fmt"

func main() {
	var nilai = 130
	var nilai32 int32 = int32(nilai)
	var nilai64 int64 = int64(nilai)
	var nilai8 int8 = int8(nilai)

	fmt.Println(nilai)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Sahal"
	var nameS = name[0]
	var nameSvalue = string(nameS)

	fmt.Println(name)
	fmt.Println(nameS)
	fmt.Println(nameSvalue)
}
