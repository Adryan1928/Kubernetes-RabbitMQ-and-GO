package main

import (
	"fmt"
)

func main() {
	numberOfBytes, errs := fmt.Println("Hello, World!", "Adryan", 100);

	fmt.Printf("Error occurred: %v\n", errs)
	if errs != nil {
		fmt.Printf("Error occurred: %v\n", errs)
	} else {
		fmt.Printf("Number of bytes: %d\n", numberOfBytes)
	}
}
