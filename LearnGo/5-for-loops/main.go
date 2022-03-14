package main

import "fmt"

func main() {

	// classic for
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// breaking from a loop
	for {
		fmt.Println("We're in the for loop")
		break
	}
	fmt.Println("Now we're out!")

	// continuing in a loop
	for n := 0; n <= 10; n++ {
		if n%3 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
