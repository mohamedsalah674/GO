package main 

import (
	// Used to Print and Scan functions
	"fmt"
)


func main() {


	myLang:= "Hello from Golang"
	fmt.Printf("%v \n " , myLang )

	fmt.Printf("Input : ")
	
	// Var declarations
	var inputString string 
	fmt.Scan(&inputString)

	fmt.Printf("%v", inputString )

}