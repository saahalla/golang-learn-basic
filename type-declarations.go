package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var nikSahal NoKtp = "340407070606566662"
	var maritalStatus Married = false

	fmt.Println(nikSahal)
	fmt.Println(maritalStatus)
}
