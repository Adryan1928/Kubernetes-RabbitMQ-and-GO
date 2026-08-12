package main

import (
	"fmt"
)

type pessoa struct {
	Nome  string
	Idade int
}

type cliente struct {
	pessoa
	estudante bool
}

func main() {
	cliente1 := cliente{
		pessoa: pessoa{
			Nome:  "João",
			Idade: 30,
		},
		estudante: true,
	}

	cliente2 := cliente{
		pessoa: pessoa{
			Nome:  "Maria",
			Idade: 25,
		},
		estudante: false,
	}

	cliente3 := struct {
		pessoa
		estudante bool
	}{
		pessoa: pessoa{
			Nome:  "Pedro",
			Idade: 35,
		},
		estudante: true,
	}

	fmt.Println("Cliente 1:", cliente1)
	fmt.Printf("Name: %s, Idade: %d, Estudante: %t\n", cliente1.Nome, cliente1.Idade, cliente1.estudante)
	fmt.Println("Cliente 2:", cliente2)
	fmt.Printf("Name: %s, Idade: %d, Estudante: %t\n", cliente2.Nome, cliente2.Idade, cliente2.estudante)
	fmt.Println("Cliente 3:", cliente3)
	fmt.Printf("Name: %s, Idade: %d, Estudante: %t\n", cliente3.Nome, cliente3.Idade, cliente3.estudante)
}