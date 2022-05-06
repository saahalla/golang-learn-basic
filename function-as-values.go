package main

import "fmt"

func getGoodNight(name string) string {
	return "good night " + name
}

func main() {
	goodNight := getGoodNight

	sahal := goodNight("Sahal")
	fmt.Println(sahal)
	fmt.Println(getGoodNight("Aput"))
}
