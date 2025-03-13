package main

import "fmt"

func main() {
	// var + nome da variável + tipo
	// nome da variável + = (recebe) + valor atribuído
	var nome string
	nome = "daniel"
	fmt.Println(nome)

	nome = "lindo"
	fmt.Println(nome)

	var idade int
	idade = 18
	fmt.Println(idade)

	// var + nome da variável + = (recebe) + valor atribuído
	var a = "futebol"
	fmt.Println(a)

	var b = true
	fmt.Println(b)

	// var + nome da variável 1 + nome da variável 2 + tipo + = (recebe) + valores atribuídos
	var c, d int = 1, 2
	fmt.Println(c)
	fmt.Println(d)

	// nome da variável + := (declarar e atribuir) + valor atribuído
	e := "bola"
	fmt.Println(e)

	f := 2006
	fmt.Println(f)

	g := 2.30
	fmt.Println(g)

	// const + nome da constante + = (recebe) + valor atribuído
	const namorada = "nenhuma"
	fmt.Println(namorada)
}