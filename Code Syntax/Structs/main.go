package main

import (
	"fmt"
)

type Cars struct {
	name  string
	color string
	price int
	model int
}

func main() {

	var car1 Cars = Cars{"name 1", "color 1", 100, 100}

	fmt.Println(car1)
	fmt.Println(car1.name)

	car1.model = 20

	fmt.Println(car1)

	// And etc ...
}
