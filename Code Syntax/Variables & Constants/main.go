package main

import (
	// Used to Print and Scan functions
	"fmt"
)

func main() {

	// I Can modify var varDef
	varDef := "X"
	fmt.Println(varDef)

	varDef = "Y"
	fmt.Println(varDef)

	// I can not modify constDef
	const constDef = "X"
	fmt.Println(constDef)

	/*
		Will cause an Error
		constDef= "E"
		fmt.Println(constDefB)
	*/

}
