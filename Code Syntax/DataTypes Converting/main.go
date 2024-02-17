package main

import (
	"fmt"
)

func main() {
	// int , int8 , int16 , int32 , int64
	// Every datatype has range
	const A int8 = 127
	const B int16 = 32767
	const C int32 = 2147483647
	const D int64 = 9223372036854775807
	fmt.Println(A, B, C, D)

	const E float32 = float32(100.22)
	fmt.Println(E)

	const F int = int(100)
	fmt.Println(F)

}
