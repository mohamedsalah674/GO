package main

import (
	"fmt"
)

func main() {
	A := "10"
	B := "11"

	if A == "10" || B == "11" {
		fmt.Println("Vaild Or")
	}

	if A == "10" && B == "11" {
		fmt.Println("Vaild And")
	}

}
