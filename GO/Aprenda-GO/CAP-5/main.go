package main

import (
	"fmt"
)

var q int = 0

const c1 = 0
const c2 int = 0

const (
	year int = iota + 2026
	year1
	year2
	year3
)

func main() {
	fmt.Println("\tWelcome")
	for {
		fmt.Println("Choose one question:")
		fmt.Scanln(&q)

		if (q == 0) {
			fmt.Println("Thank you for use my program")
				break
		}

		switch (q) {
			case 1:
				Q1()
			case 2:
				Q2()
			case 3:
				Q3()
			case 4:
				Q4()
			case 5:
				Q5()
			case 6:
				Q6()
			default:
				fmt.Println("Invalid value")
		}
	}
}

func Q1() {
	fmt.Println("Ex - 01")
	x := 10

	fmt.Printf("%v, %b, %X \n", x, x, x)

}

func Q2() {
	fmt.Println("Ex - 02")
	x := 10
	y := 12

	fmt.Println(x > y, " - ", x >= y, " - ", x < y, " - ", x <= y, " - ", x != y, " - ", x == y)
}

func Q3() {
	fmt.Println("Ex - 03")

	fmt.Printf("%v, %T - %v, %T\n", c1, c1, c2, c2)
}

func Q4() {
	fmt.Println("Ex - 04")

	x := 10

	fmt.Printf("%v, %b, %X \n", x, x, x)

	y := x << 1

	fmt.Printf("%v, %b, %X \n", y, y, y)
}

func Q5() {
	fmt.Println("Ex - 05")

	x := ` oi
		vc
					está

ai?????
	`

	fmt.Println(x)
}

func Q6() {
	fmt.Println("Ex - 05")

	fmt.Printf("%v, %v, %v, %v\n", year, year1, year2, year3)
}