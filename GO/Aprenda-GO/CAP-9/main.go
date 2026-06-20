package main

import (
	"fmt"
)

var q int = 0

func main() {
	fmt.Println("Welcome to my Go program!")

	for {
		fmt.Println("Choose a question:")
		fmt.Scanln(&q)

		if q == 0 {
			fmt.Println("Exiting the program. Goodbye!")
			break
		}

		switch q {
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
				fmt.Println("Invalid question number. Please try again.")
		}
	}
}

func Q1() {
	fmt.Println("Question 1:")
	array := [5]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("Array type: %T\n", array)
}


func Q2() {
	fmt.Println("Question 2:")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("Slice type: %T\n", slice)
}


func Q3() {
	fmt.Println("Question 3:")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:len(slice)-1])
}

func Q4() {
	fmt.Println("Question 4:")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)
}

func Q5() {
	fmt.Println("Question 5:")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := x[:3]
	y = append(y, x[6:]...)
	fmt.Println(y)
}

func Q6() {
	fmt.Println("Question 6:")
	slice := make([]string, 26)
	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, slice[i])
	}

}

func Q7() {
	fmt.Println("Question 7:")
	slice := [][]string{
		{"Nome", "Sobrenome", "Hobby"},
		{"João", "Silva", "Futebol"},
		{"Maria", "Souza", "Leitura"},
		{"Pedro", "Oliveira", "Música"},
	}

	for i, v := range slice {
		for j, val := range v {
			fmt.Printf("Row: %d, Column: %d, Value: %s\n", i, j, val)
		}
	}
}

func Q8() {
	fmt.Println("Question 8:")

	people := map[string][]string{
		"João": {"Basquete", "Futebol"},
		"Maria": {"Leitura", "Cinema"},
		"Pedro": {"Música", "Viagem"},
	}

	for name, hobbies := range people {
		fmt.Printf("Name: %s\n", name)
		for i, hobby := range hobbies {
			fmt.Printf("\tHobby %d: %s\n", i+1, hobby)
		}
	}
}

func Q9() {
	fmt.Println("Question 9:")

	people := map[string][]string{
		"João": {"Basquete", "Futebol"},
		"Maria": {"Leitura", "Cinema"},
		"Pedro": {"Música", "Viagem"},
	}

	people["Ana"] = []string{"Dança", "Culinária"}

	for name, hobbies := range people {
		fmt.Printf("Name: %s\n", name)
		for i, hobby := range hobbies {
			fmt.Printf("\tHobby %d: %s\n", i+1, hobby)
		}
	}
}

func Q10() {
	fmt.Println("Question 10:")

	people := map[string][]string{
		"João": {"Basquete", "Futebol"},
		"Maria": {"Leitura", "Cinema"},
		"Pedro": {"Música", "Viagem"},
	}

	people["Ana"] = []string{"Dança", "Culinária"}

	delete(people, "Maria")

	for name, hobbies := range people {
		fmt.Printf("Name: %s\n", name)
		for i, hobby := range hobbies {
			fmt.Printf("\tHobby %d: %s\n", i+1, hobby)
		}
	}
}