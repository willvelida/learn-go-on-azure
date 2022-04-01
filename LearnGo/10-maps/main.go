package main

import (
	"fmt"
)

func main() {
	var crusaders = map[string]string{"Team": "Crusaders", "Location": "Christchurch"}
	var blues = make(map[string]string)
	blues["Team"] = "Blues"
	blues["Location"] = "Wellington"

	fmt.Println(crusaders)
	fmt.Println(blues)

	fmt.Println(blues["Location"])

	blues["Color"] = "Blue"
	crusaders["Color"] = "Red"

	fmt.Println(crusaders)
	fmt.Println(blues)

	delete(blues, "Location")
	blues["Location"] = "Auckland"
	fmt.Println(blues["Location"])

	for k, v := range crusaders {
		fmt.Println(k, v)
	}

	for _, element := range blues {
		fmt.Println(element, blues[element])
	}
}
