package main

import (
	"fmt"
)

type myType int

var z int = 5

func main() {
	x := 10
	y := "String exemple"

	fmt.Printf("X: %v, %T \n", x, x)
	fmt.Printf("Y: %v, %T \n", y, y)

	x = 20

	fmt.Printf("X: %v, %T \n", x, x)

	x, z = swap(x, z)

	fmt.Printf("X: %v, %T \n", x, x)
	fmt.Printf("Z: %v, %T \n", z, z)
}

func swap(a int, b int) (int, int) {
	return b, a
}