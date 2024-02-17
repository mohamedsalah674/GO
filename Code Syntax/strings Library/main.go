package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	A := "Hello"
	B := "hello"

	// Returns -1 because A != B
	fmt.Println(strings.Compare(A, B))

	// Returns true because A contains H
	fmt.Println(strings.Contains(A, "H"))

	// Returns true because A contains H
	fmt.Println(strings.ContainsAny(A, "111H111"))

	// Returns 2 because A contains two l letters
	fmt.Println(strings.Count(A, "l"))

	// Returns Before and After and Found
	// Found --> true  --> bol
	// Before --> He   --> bef
	// After --> o	   --> af
	// Because ll found in A
	af, bef, bol := strings.Cut(A, "ll")
	fmt.Println(af, bef, bol)

	// Compare and not Case Sensitivie
	// Returns true
	// Works only for strings
	fmt.Println(strings.EqualFold(A, B))

	// Returns sclice include strings of parameters of Field
	slice := strings.Fields("firstElement lastElement")
	fmt.Println(slice)

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)

	}
	fmt.Println(strings.FieldsFunc("elm1 elm2 elm3 ", f))
}
