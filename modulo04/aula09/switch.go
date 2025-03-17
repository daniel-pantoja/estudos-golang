package main

import "fmt"

func main() {
	posicao := 3
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "Programação"
	fmt.Println(nome)
	switch nome {
	case "Daniel":
		fmt.Println("É meu nome")
	case "Futebol":
		fmt.Println("É o esporte favorito de Daniel")
	case "Programação":
		fmt.Println("É o foco!")
	}
}