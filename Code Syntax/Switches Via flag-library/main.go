package main

import (
	"flag"
	"fmt"
)

// flag.dataType
// flag.string(	"name", "Mohamed Salah" , "This is my name")
// flag.int("age", 100 , "This is age ")

func main() {

	name := flag.String("name", "Mohamed Salah", "This is my name")
	age := flag.Int("age", 24, "This is age ")

	// Print the address
	fmt.Println(name, age)

	// Print Values
	fmt.Println(*name, *age)
}
