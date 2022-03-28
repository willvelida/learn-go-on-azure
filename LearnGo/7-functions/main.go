package main

import "fmt"

func main() {
	greeting()
	greetAPerson("Will")
	fullName := constructName("Will", "Velida")
	greetAPerson(fullName)
	name, futureAge := ageInFiveYears(fullName, 31)
	fmt.Println(name, futureAge)
	_, unknownPersonAge := ageInFiveYears(fullName, 40)
	fmt.Println(unknownPersonAge)
	countToFive(1)
}

func greeting() {
	fmt.Println("Hello there!")
}

func greetAPerson(name string) {
	fmt.Println("Hello", name)
}

func constructName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func ageInFiveYears(name string, ageNow int) (greeting string, ageInFiveYears int) {
	greeting = "My name is " + name + " and in 5 years, I will be"
	ageInFiveYears = ageNow + 5
	return
}

func countToFive(x int) int {
	if x == 6 {
		return 0
	}
	fmt.Println(x)
	return countToFive(x + 1)
}
