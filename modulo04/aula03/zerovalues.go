package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool 
	var s string = "Daniel"
	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %v\n", s)

	i = 15000
	fmt.Printf("Inteiro: %v\n", i)

	f = 1.5
 	fmt.Printf("Float: %v\n", f)

	b = true
	fmt.Printf("Bool: %v\n", b)

	s = "Pantoja"
	fmt.Printf("String: %v\n", s)
}