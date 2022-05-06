package main

import "fmt"

type Todo struct {
	ID         int
	ActivityID int
	Title      string
}

func main() {
	todo1 := Todo{
		ID:         1,
		ActivityID: 1,
		Title:      "todo items 1",
	}
	todo1.ID = 2
	fmt.Println(todo1)

	todo2 := Todo{
		ID:         2,
		ActivityID: 1,
		Title:      "todo-items 2",
	}
	fmt.Println(todo2)
}
