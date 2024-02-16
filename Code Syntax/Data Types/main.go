package main

import (
	// Used to Print and Scan functions
	"fmt"
)

func main() {

	// Boolean
	myBoolean := false
	fmt.Printf("%v \n ", myBoolean)

	// Integer
	myInteger := 100
	fmt.Printf("%v \n ", myInteger)

	// Float
	myFloat := 100.2
	fmt.Printf("%v \n ", myFloat)

	// String
	myString := "Hello from Golang"
	fmt.Printf("%v \n ", myString)

	// Array

	var myStringArray = []string{}
	myStringArray = append(myStringArray, "Hello from Golang 1 ")
	myStringArray = append(myStringArray, "Hello from Golang 2 ")

	fmt.Printf("%v \n ", myStringArray)

}
