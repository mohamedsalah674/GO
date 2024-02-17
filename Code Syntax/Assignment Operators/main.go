package main

import (
	"fmt"
)

func main() {

	// = , == , += , -= , *= , /=

	var A = 10
	var B = 20

	// A == B ---> Returns Boolean value
	if A == B {
		fmt.Println("A is equal to B")
	}

	// A!= B ---> Returns Boolean value
	if A != B {
		fmt.Println("A is not equal to B")
	}

	// A > B ---> Returns Boolean value
	if A > B {
		fmt.Println("A is greater than B")
	}

	// A >= B ---> Returns Boolean value
	if A >= B {
		fmt.Println("A is greater than or equal to B")
	}

	// A < B ---> Returns Boolean value
	if A < B {
		fmt.Println("A is less than B")
	}

	// A <= B ---> Returns Boolean value
	if A <= B {
		fmt.Println("A is less than or equal to B")
	}

	// A += B ---> A = A + B
	A += B
	fmt.Println(A)

	// A -= B ---> A = A - B
	A -= B
	fmt.Println(A)

	// A *= B ---> A = A * B
	A *= B
	fmt.Println(A)

	// A /= B ---> A = A / B
	A /= B
	fmt.Println(A)

	// A % B ---> A = A % B
	A %= B
	fmt.Println(A)

}
