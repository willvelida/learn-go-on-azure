package main

import "fmt"

func main() {
	number1 := 10
	number2 := 20

	// if, else
	if number1 > number2 {
		fmt.Println("10 is bigger than 20")
	} else {
		fmt.Println("No, 10 is not bigger than 20.")
	}

	// if, else if
	if 9%2 == 0 {
		fmt.Println("9 is a even number")
	} else if 9 > 10 {
		fmt.Println("9 is also bigger than 10!")
	} else {
		fmt.Println("No, 9 isn't a even number nor is it bigger than 10.")
	}

	// Ifception
	if number1 > 9 {
		if number2 > 19 {
			fmt.Println("Both conditions are true")
		} else {
			fmt.Println("Only the first condition is true")
		}
	} else {
		fmt.Println("The first condition is false")
	}

	// Conditions
	fmt.Println(number1 < number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 <= number2)
	fmt.Println(number1 >= number2)
	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
}
