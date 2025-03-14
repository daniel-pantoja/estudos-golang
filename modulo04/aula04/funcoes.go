package main

import "fmt"

func main() {
	fmt.Println(soma(11, 10))

	sub := subtracao(15, 15)
	fmt.Println(sub) 

	mult := multiplicacao(2, 2)
	fmt.Println(mult)

	mais, menos := somaSubtracao(2000, 6, 30, 12)
	fmt.Println(mais)
	fmt.Println(menos)

	fmt.Println(hobby("Futebol"))

	nome, sobrenome := printNomeSobrenome("Daniel", "Pantoja")
	fmt.Println(nome)
	fmt.Println(sobrenome)

	sport1, sport2, sport3 := esporte("Futsal", "Futebol", "Fut7")
	fmt.Println(sport1)
	fmt.Println(sport2)
	fmt.Println(sport3)
}

func soma(x int, y int, ) int {
	return x + y
}

func subtracao (z, w int) int {
	return z - w
}

func multiplicacao(x, y int) int {
	return x * y
}

func somaSubtracao(x, y, z, w int) (int, int) {
	return x + y, z - w
}

func hobby(hobby string) string {
	return hobby
}

func printNomeSobrenome(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

func esporte(sport1, sport2, sport3 string) (string, string, string) {
	return sport1, sport2, sport3
} 


// Função começando com letra minúscula --> função PRIVADA
// Função começando com letra maiúscula --> função PÚBLICA