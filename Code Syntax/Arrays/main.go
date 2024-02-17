package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	var arr2 [3]string
	arr2[0] = "one"
	arr2[1] = "two"
	arr2[2] = "three"
	fmt.Println(arr2)

	for _, x := range arr {
		if x == arr[x-1] {
			fmt.Println(x)
		}

	}

}
