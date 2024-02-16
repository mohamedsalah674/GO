package main

import (
	// Used to Print and Scan functions
	"fmt"
)

func main() {

	fmt.Println("Enter Your Name: ")

	// Take input
	var userName string
	fmt.Scanln(&userName)

	fmt.Println(userName)

}
