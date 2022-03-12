package main

import "fmt"

func main() {
	number1 := 10
	number2 := 20

	if number1 > number2 {
		fmt.Println("10 is bigger than 20")
	} else {
		fmt.Println("No, 10 is not bigger than 20.")
	}

	if 9 % 2 == 0 {
		fmt.Println("9 is a even number")
	} else if 9 > 10 {
		fmt.Println("9 is also bigger than 10!")
	} else {
		fmt.Println("No, 9 isn't a even number or is it bigger than 10.")
	}
}
