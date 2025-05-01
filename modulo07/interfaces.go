package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

type Animal interface {
	Barulho()
	NumeroDePatas()
}

func main() {
	cachorro := Cachorro{
		Raca:  "Pinscher",
		Cor:   "Preto",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "Amarelo",
		Patas: 4,
	}

	fmt.Println(cachorro.Barulho())
	fmt.Println(gato.Barulho())
}