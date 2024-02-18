package main

import (
	"fmt"
)

func sum(a, b int) int {
	c := a + b
	return c
}

func sumMulBy2(c int) int {
	d := c * 2
	return d
}

func printHello() {
	fmt.Println("Hello")
}
func main() {

	printHello()
	fmt.Println(sum(1, 2))
	fmt.Println(sumMulBy2(sum(1, 2)))
}
