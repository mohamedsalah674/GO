package main

import (
	"fmt"
)

func main() {

	A := "10"

	if A == "test" {
		fmt.Println("test")
	} else if A == "10" {
		fmt.Println("10")
	} else {
		fmt.Println(A)
	}

}
