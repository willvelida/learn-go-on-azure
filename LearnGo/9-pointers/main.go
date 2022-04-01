package main

import "fmt"

func main() {

	b := 20
	pointerToB := &b

	fmt.Println(pointerToB)  // prints memory address
	fmt.Println(*pointerToB) // prints 10

	c := 5 * *pointerToB
	fmt.Println(c) // prints 100
}
