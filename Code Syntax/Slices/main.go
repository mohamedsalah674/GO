package main

import "fmt"

func main() {
	var Slc []string = []string{"1", "2", "3", "4", "5"}
	fmt.Println(Slc)

	Slc = append(Slc, "6")
	fmt.Println(Slc)

	// And etc...
}
