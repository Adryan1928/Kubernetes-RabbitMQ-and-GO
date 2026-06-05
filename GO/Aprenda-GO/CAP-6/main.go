package main

import (
	"fmt"
)

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
}
