package main

import "fmt"

// IF = Se
// ELSE = Senão

func main() {
	
	numero := 8
// if (condição) { <ação> }
	if numero == 1 { // true
		fmt.Println("Valor é igual a 5")
	} else {
		fmt.Println("Valor Não é 5")
	}

	if numero == 5 { // true
		fmt.Println("Valor é igual a 5")
	} else if numero == 10 {
		fmt.Println("Valor é igual a 10")
	} else { // false
		fmt.Println("Valor é diferente de 5 e 10")
	}

	if 7%3 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	} 

	if numero % 2 == 0 {
		fmt.Printf("%d é par\n", numero)
	} else {
		fmt.Printf("%d é impar\n", numero)
	}

	nome := "Ramilly" 
	if nome == "Daniel" {
		fmt.Println("Daniel quer se tornar Dev!")
	} else {
		fmt.Println("Olá,", nome)
	}

	// Operações
	fmt.Println(46 + 92)
	fmt.Println(55 - 34)
	fmt.Println(2 * 3)
	fmt.Println(10 / 2)
	fmt.Println(10 % 3) // % È o resto da divisão 

	numero2 := 8
	if numero2%2 == 0 {
		fmt.Printf("%d é par\n", numero2)
	} else {
		fmt.Printf("%d é impar\n", numero2)
	}

	usuario := "Junior"
	if usuario == "Junior" {
		fmt.Println("Usuário logado!")
	} else {
		fmt.Println("Tente novamente um usuário válido.")
	}
}