package main

import "fmt"

// type + nome da estrutura + struct + { Campos }
type Pessoa struct {
	Nome string
	Idade int
}

type InformacoesPessoas struct {
	Nome string
	Sobrenome string
	Idade int
	AnoNasc int
	Altura float64
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Daniel", 18})
	fmt.Println(Pessoa{Nome: "Ramilly", Idade: 19})
	fmt.Println(Pessoa{Nome: "Daniel"})
	fmt.Println(Pessoa{Idade: 18})

	p1 := Pessoa{Nome: "Neymar", Idade: 33}
	fmt.Println(p1)
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)
	p1.Idade = 19
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Daniel-Pantoja", Idade: 18}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	infp := InformacoesPessoas{Nome: "Antenor", Sobrenome: "Neto", Idade: 21, AnoNasc: 2003, Altura: 1.64}
	fmt.Println(infp)
	fmt.Println(infp.Altura)
	fmt.Println(infp.AnoNasc)
	fmt.Println(infp.Idade)
	fmt.Println(infp.Sobrenome)
	fmt.Println(infp.Nome)

	// structs + map
	alunosTi := map[string] []Pessoa{}
	alunosTi["programação"] = pessoas
	fmt.Println(alunosTi)

	var alunosCursos = map[string] []Pessoa{
		"Programação": {{Nome: "Daniel", Idade: 18}, {Nome: "Guanabara", Idade: 52}},
		"Engenharia": {{Nome: "Taian", Idade: 22}, {Nome: "Thalis", Idade: 22 }}, 
	} 
	fmt.Println(alunosCursos)

	// struct herdando campos de outra struct 
	prof := Profissao{p2, "Dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}