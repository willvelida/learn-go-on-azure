package main

import "fmt"

type Todo struct {
	id         int
	title      string
	isComplete bool
}

func main() {
	var todoItem1 Todo

	todoItem1.id = 1
	todoItem1.title = "Write code for Struct Blog"
	todoItem1.isComplete = false

	printTodoItem(todoItem1)

	todoItem1.isComplete = true
	printTodoItem(todoItem1)

	fmt.Println(Todo{title: "Take dog for a walk", isComplete: false})
}

func printTodoItem(todoItem Todo) {
	fmt.Println("Id:", todoItem.id)
	fmt.Println("Title:", todoItem.title)
	fmt.Println("IsComplete:", todoItem.isComplete)
}
