package main

import (
	"fmt"
)

var x int = 0
var y string = "James Bond"
var z bool = true

type myType int

var a myType

var b int

func main() {
	fmt.Println("Welcome to my program!")
	for {
		fmt.Println("Choose an exercise:")
		fmt.Scanln(&x)

		if (x == 0){
			fmt.Println("Thank you for use my program!")
			break
		}

		switch (x){
			case (1):
				Q1()
			case (2):
				Q2()
			case (3):
				Q3()
			case (4):
				Q4(true)
			case (5):
				Q5()
			default:
				fmt.Println("Invalid Value")
		}
	}
}


func Q1() {
	fmt.Println("Queston one:")
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("X: %v, Y: %v and Z: %v \n", x, y, z)

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)
}


func Q2() {
	fmt.Println("Question two")

	fmt.Printf("Value: %v,  type: %T \n", x, x)
	fmt.Printf("Value: %v,  type: %T \n", y, y)
	fmt.Printf("Value: %v,  type: %T \n", z, z)
}


func Q3() {
	fmt.Println("Question three")

	s := fmt.Sprintf("%v%v%v", x, y, z)

	fmt.Println(s)
}

func Q4(showPrint bool) {
	if (showPrint){
		fmt.Println("Question four")
	}

	fmt.Printf("valor: %v, Type: %T\n", a, a)

	a = 42

	fmt.Printf("valor: %v, Type: %T\n", a, a)
}

func Q5() {
	fmt.Println("Question five")

	Q4(false)

	b = int(a)

	fmt.Printf("Valor: %v, type: %T\n", b, b)

}