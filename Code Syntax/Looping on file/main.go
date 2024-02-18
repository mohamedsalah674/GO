package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(scanner)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
