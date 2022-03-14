package main

import "fmt"

func main() {
	favouriteColor := "red"

	switch favouriteColor {
	case "blue":
		fmt.Println("My favourite color is blue")
	case "green":
		fmt.Println("My favorite color is green")
	case "red":
		fmt.Println("My favorite color is red")
	default:
		fmt.Println("I don't have a favorite color")
	}

	var number int
	fmt.Print("Please enter a number between 1 and 9: ")
	fmt.Scan(&number)

	switch {
	case number%2 == 0:
		fmt.Println("The number you entered is even")
	default:
		fmt.Println("The number you entered is odd")
	}
}
