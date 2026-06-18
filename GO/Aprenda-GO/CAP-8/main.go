package main

import (
	"fmt"
)

var array [5]int
var slice []int
var x = [5]int{1, 2, 3, 4, 5}

var matrix = [][]int{
	{1, 2, 3},
	{4, 5, 6},
}

var dictionary = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
}

func main() {
	x[0] = 42
	fmt.Println(x[0])
	fmt.Println(x)
	fmt.Println(len(x))

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))

	stringSlice := []string{"Go", "is", "awesome"}
	for index, value := range stringSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fatia := stringSlice[:2]
	fmt.Println(fatia)

	fatia2 := append(stringSlice[:1], stringSlice[2:]...)
	fmt.Println(fatia2)


	sliceWithMake := make([]int, 3, 5)
	sliceWithMake[0], sliceWithMake[1], sliceWithMake[2] = 10, 20, 30
	fmt.Println(sliceWithMake)
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceWithMake), cap(sliceWithMake))

	sliceWithMake = append(sliceWithMake, 40, 50)
	fmt.Println(sliceWithMake)
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceWithMake), cap(sliceWithMake))

	sliceWithMake = append(sliceWithMake, 60)
	fmt.Println(sliceWithMake)
	fmt.Printf("Length: %d, Capacity: %d\n", len(sliceWithMake), cap(sliceWithMake))


	fmt.Println(matrix)

	fmt.Println(dictionary)

	value, ok := dictionary["four"]
	fmt.Printf("Value: %d, Exists: %t\n", value, ok)
}