package main

import (
	// Used to Print and Scan functions
	"fmt"
)

func main() {

	// Method 1
	var myVar1 string = "This is my variable 1"
	fmt.Println(myVar1)

	// Method 2
	var myVar2 string
	myVar2 = "This is my variable 2"
	fmt.Println(myVar2)

	// Without Type
	myVar3 := ""
	fmt.Println(myVar3)

	// Define Multiple Variables
	myVar4, myVar5 := "This is my variable 4 ", "This is my variable 5 "
	fmt.Print(myVar4, " \n", myVar5, " \n")

	// Define Multiple Variables
	var myVar6, myVar7 = "This is my variable 6 ", "This is my variable 7 "
	fmt.Print(myVar6, " \n", myVar7, " \n")

	// Define Multiple Variables intgers
	var a, b = 1, 2
	fmt.Print(a, " \n", b, " \n")

	var z, y int
	z = 5
	y = 3

	fmt.Print(z, " \n", y, " \n")

	// Define Slice
	mySlice1 := []string{"A", "B"}
	fmt.Println(mySlice1)

	var mySlice2 = []string{"C", "D"}
	fmt.Println(mySlice2)

}
