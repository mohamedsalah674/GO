package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}

	for i := 10; i > 0; i-- {
		if i == 5 {
			fmt.Println("Inside Loop")
			continue
		}

		if i == 1 {
			break
		}

		fmt.Println(i)

	}
}
