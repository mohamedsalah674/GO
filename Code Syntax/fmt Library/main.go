package main

import "fmt"

func main() {

	// For Strings

	A := "DEF"
	println("ABC" + A)

	// Print Value of Variable
	fmt.Printf("%v", A)
	print("\n")

	// Print Type of Variable
	fmt.Printf("%T", A)
	print("\n")

	// Print String of Variable
	fmt.Printf("%s", A)
	print("\n")

	// Print Base 16 of Variable of small characters
	fmt.Printf("%x", A)
	print("\n")

	// Print Base 16 of Variable of large characters
	fmt.Printf("%X", A)
	print("\n")

	// For Integers

	B := 15

	// Print 15
	fmt.Printf("%c", B)
	fmt.Println("")

	// base 10
	fmt.Printf("%d", B)
	fmt.Println("")

	// Base 8
	fmt.Printf("%o", B)
	fmt.Println("")

	// Single Quote Character
	fmt.Printf("%q", B)
	fmt.Println("")

	// Base 16 of Variable of small characters
	fmt.Printf("%x", B)
	fmt.Println("")

	// Base 16 of Variable of large characters
	fmt.Printf("%X", B)
	fmt.Println("")

	// Unicode Format
	fmt.Printf("%U", B)
	fmt.Println("")

	// Built in GO without need to import fmt
	/*
		A := "DEF"
		print("ABC" + A)
		print("\n")
	*/
}
