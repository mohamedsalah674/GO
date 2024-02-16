package main

import (
	// Used to Print and Scan functions
	"fmt"
)

func main() {

	//     \b --> Back Space
	// Expected output -- > Helloworld
	fmt.Println("Hello \bworld!")

	//     \n --> New Line
	// Expected output --> Hello
	//					   world
	fmt.Println("Hello \nworld!")

	// \r --> return
	fmt.Println("AAAA\rFour")

	// \t --> Tab
	fmt.Println("\t Hello World!")

	// Escape Character
	// Expected output --> Hello/World!
	fmt.Println("Hello \\ World!")

	// use the escape character --> ``
	// Expected output --> Hello
	//					   World!
	//					   World!
	fmt.Println(`Hello
World!
World!
	`)

}
