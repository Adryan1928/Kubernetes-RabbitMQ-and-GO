package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my program!")
	q := 1

	for q != 0 {
		fmt.Println("Choose a question:")
		fmt.Scanln(&q)

		switch q {
			case 1:
				Q1()
			case 2:
				Q2()
			case 3:
				Q3()
			case 4:
				Q4()
			default:
				if q != 0 {
					fmt.Println("Invalid question number. Please try again.")
				}
		}
	}

	fmt.Println("Thanks for use my program")
}

func Q1() {
	fmt.Println("Question 1:")

	type Pessoa struct {
		nome  string
		sobrenome string
		sabores []string
	}

	p1 := Pessoa{
		nome: "João",
		sobrenome: "Silva",
		sabores: []string{"Chocolate", "Baunilha", "Morango"},
	}

	p2 := Pessoa{
		nome: "Maria",
		sobrenome: "Souza",
		sabores: []string{"Menta", "Coco", "Caramelo"},
	}
	
	fmt.Println("Pessoa 1:", p1)
	fmt.Printf("Nome: %s\nSobrenome: %s\n", p1.nome, p1.sobrenome)
	fmt.Println("Sabores favoritos:")
	for _, sabor := range p1.sabores {
		fmt.Println("-", sabor)
	}

	fmt.Println("\nPessoa 2:", p2)
	fmt.Printf("Nome: %s\nSobrenome: %s\n", p2.nome, p2.sobrenome)
	fmt.Println("Sabores favoritos:")
	for _, sabor := range p2.sabores {
		fmt.Println("-", sabor)
	}
}

func Q2() {
	fmt.Println("Question 2:")

	type Pessoa struct {
		nome  string
		sobrenome string
		sabores []string
	}

	dict := make(map[string]Pessoa)

	dict["silva"] = Pessoa{
		nome: "João",
		sobrenome: "Silva",
		sabores: []string{"Chocolate", "Baunilha", "Morango"},
	}

	dict["souza"] = Pessoa{
		nome: "Maria",
		sobrenome: "Souza",
		sabores: []string{"Menta", "Coco", "Caramelo"},
	}

	for key, pessoa := range dict {
		fmt.Printf("Key: %s\n", key)
		fmt.Printf("Nome: %s\nSobrenome: %s\n", pessoa.nome, pessoa.sobrenome)
		fmt.Println("Sabores favoritos:")
		for _, sabor := range pessoa.sabores {
			fmt.Println("-", sabor)
		}
		fmt.Println()
	}
}

func Q3() {
	fmt.Println("Question 3:")
	
	type veiculo struct {
		portas int
		cor string
	}

	type caminhonete struct {
		veiculo
		quatroRodas bool
	}

	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	c1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor: "Preto",
		},
		quatroRodas: true,
	}

	s1 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor: "Branco",
		},
		modeloLuxo: true,
	}

	fmt.Println("Caminhonete:", c1)
	fmt.Printf("Portas: %d, Cor: %s, Quatro Rodas: %t\n", c1.portas, c1.cor, c1.quatroRodas)
	fmt.Println("Sedan:", s1)
	fmt.Printf("Portas: %d, Cor: %s, Modelo Luxo: %t\n", s1.portas, s1.cor, s1.modeloLuxo)
}

func Q4() {
	fmt.Println("Question 4:")

	type anonimo struct {
		data map[string]string
		lista []string
	}

	a1 := anonimo{
		data: map[string]string{
			"chave1": "valor1",
			"chave2": "valor2",
		},
		lista: []string{"item1", "item2", "item3"},
	}

	fmt.Println("Anonimo:", a1)
	fmt.Println("Data:", a1.data)
	fmt.Println("Lista:", a1.lista)
}