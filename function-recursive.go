// factorial loop

package main

import "fmt"

func factorialLoop(value int) (result int) {
	result = 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fact20 := factorialLoop(20)
	factr20 := factorialRecursive(20)
	fmt.Println(fact20, factr20)
}
