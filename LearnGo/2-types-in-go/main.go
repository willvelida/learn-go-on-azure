package main

import "fmt"

const statement string = "Learning Go is Fun!"

func main() {

	// strings
	var greeting string = "This is a string!"
	fmt.Println(greeting)

	// booleans
	var falseValue bool
	var isBooleanTrue = true
	fmt.Println(falseValue)
	fmt.Println(isBooleanTrue)

	// Integers and floats
	number1 := 5
	number2 := 3
	fmt.Println("number1 + number2 = ", number1+number2)
	var number3 float64 = 5.0
	number4 := 3.0
	fmt.Println("number3 + number4 = ", number3+number4)

	fmt.Println(statement)

}
