package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"Cat":  50,
		"Dog":  60,
		"Bird": 90,
	}

	fmt.Println(myMap)

	// Can I Modify this
	myMap["Cat"] = 60
	fmt.Println(myMap)

	// Can I delete this
	delete(myMap, "Cat")
	fmt.Println(myMap)
}
