package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
}

type gente interface {
	falar()
}

func (p pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.nome, "e tenho", p.idade, "anos")
}

type dentista struct {
	pessoa
	dentes int
}

type engenheiro struct {
	pessoa
	plantas int
}

func serhumano(g gente) {
	g.falar()
}

func main() {
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(soma(si...))

	defer fmt.Println("last")
	fmt.Println("first")

	pessoa1 := pessoa{"João", 30}
	pessoa1.falar()

	dentista1 := dentista{pessoa{"Maria", 28}, 32}
	dentista1.falar()
	fmt.Println("Número de dentes:", dentista1.dentes)

	engenheiro1 := engenheiro{pessoa{"Carlos", 35}, 10}
	engenheiro1.falar()
	fmt.Println("Número de plantas:", engenheiro1.plantas)

	serhumano(dentista1)
	serhumano(engenheiro1)

	func (x int) {
		fmt.Println("Função anonima executada com valor:", x)
	}(42)

	newSoma := somax10()
	fmt.Println(newSoma(1, 2, 3))

	fmt.Println(somentepares(soma, []int{1,2,3,4,5,6,7,8,9,10}...))
}

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x)
}

func somax10() func(...int) (int, int) {

	return func(y ...int) (int, int) {
		soma := 0
		for _, v := range y {
			soma += v
		}
		return soma + 10, len(y)
	}
}

func somentepares(f func(x ...int) (int, int), x ...int) int {
	var slice []int
	soma := 0
	for _, v := range x {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	soma, _ = f(slice...)

	return soma
}