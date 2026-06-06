package main

import (
	"fmt"
)

var x interface{}

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %U - %v \n", x, x, string(x))
	}

	if x := 10; x == 10 {
		fmt.Println("True")
	} else if (x > 100){
		fmt.Println("True tbm")
	} else {
		fmt.Println("false")
	}

	x = 10

	switch x.(type) {
		case int:
			fmt.Println("É um Int!")
		case bool:
			fmt.Println("É um Booleano!")
		case string:
			fmt.Println("É uma String!")
		case float64:
			fmt.Println("É um Float!")
		default:
			fmt.Println("Não tem tipo")
	}
}
