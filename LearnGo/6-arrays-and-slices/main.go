package main

import "fmt"

func main() {

	// declaring an array using var
	var myFavoriteNumbers = [3]int{8, 7, 19}

	// declaring an array using shorthand
	myFavoriteColors := [4]string{"Red", "Blue", "Green", "Orange"}

	fmt.Println(myFavoriteNumbers)

	feelingKinda := myFavoriteColors[1]
	fmt.Println("I'm feeling kinda " + feelingKinda + " today")

	// Creating an empty array
	var emptyArray [3]int
	// Setting a value in that int
	fmt.Println(emptyArray)
	emptyArray[1] = 10
	fmt.Println(emptyArray)
	emptyArray[0] = 5
	fmt.Println(emptyArray)
	emptyArray[2] = 9
	fmt.Println(emptyArray)

}
