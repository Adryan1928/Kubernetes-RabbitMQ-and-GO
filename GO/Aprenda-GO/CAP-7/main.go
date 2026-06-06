package main

import (
	"fmt"
	"time"
)

func main() {
	q := 1
	for q != 0 {
		fmt.Println("Choose one question: ")
		fmt.Scanln(&q)

		switch (q) {
			case 0:
				fmt.Println("Thank you for use my program!")
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
			case 7:
				Q7()
			case 8:
				Q8()
			case 9:
				Q9()
			case 10:
				Q10()
			default:
				fmt.Println("Invalid value")
		}
	}
}

func Q1() {
	fmt.Println("EX-01")
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func Q2() {
	fmt.Println("EX-02")

	word := 65

	for i := 0; i < 26; i++ {
		fmt.Println(word + i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", word + i)
		}
		
	}
}

func Q3() {
	fmt.Println("EX-03")

	for year := 2006; year <= time.Now().Year(); year++ {
		fmt.Println(year)
	}
}

func Q4() {
	fmt.Println("EX-04")

	year := 2006
	for {
		if (year > time.Now().Year()){
			break
		}
		fmt.Println(year)
		year++
	}
}

func Q5() {
	fmt.Println("EX-05")

	for x := 10; x <= 100; x++ {
		fmt.Printf("%v %% 4 == %v\n", x, x % 4)
	}
}

func Q6() {
	fmt.Println("EX-06")

	x := 10

	if x == 10 {
		fmt.Println("X equal ten")
	}
}

func Q7() {
	fmt.Println("EX-07")

	x := 10

	if x == 10 {
		fmt.Println("X equal ten")
	} else if (x < 10) {
		fmt.Println("x lower than ten")
	} else {
		fmt.Println("x bigger than ten")
	}
}

func Q8() {
	fmt.Println("EX-08")

	x := 10

	switch {
		case x == 10:
			fmt.Println("X equal ten")
		case x < 10:
			fmt.Println("x lower than ten")
		default:
			fmt.Println("x bigger than ten")
	}
}

func Q9() {
	fmt.Println("EX-09")

	esporteFavorito := "futebol"

	switch esporteFavorito {
		case "futebol":
			fmt.Println("Seu esport favorito é Futebol!")
		case "volei":
			fmt.Println("Seu esport favorito é Volei!")
		case "basquete":
			fmt.Println("Seu esport favorito é Basquete!")
	}
}

func Q10() {
	fmt.Println("EX-10")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}