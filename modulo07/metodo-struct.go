// Atrelar o método a um tipo específico de struct
// Quando altero o método a uma struct, eu tenho acesso a todos os campos desta struct

package main

import "fmt"
type Pessoa struct {
	Nome string
	Idade int
	Profissao string
	Hobby string
	Altura float64
	Comida string
}

func (p Pessoa) ListaNomeEIdade() string {
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func (p2 Pessoa) ListaNomeEIdade2() string {
	return fmt.Sprintf("%s tem %d anos e seu Hobby é %s", p2.Nome, p2.Idade, p2.Hobby)
}

func (ds Pessoa) DescricaoDaPessoa() string {
	return fmt.Sprintf("%s é %s tem %d anos, %f de altura sua comida favorita é %s e seu hobby é %s", ds.Nome, ds.Profissao, ds.Idade, ds.Altura, ds.Comida, ds.Hobby)
}

func main() {
	
	pessoa := Pessoa {
		Nome: "Daniel Jr",
		Idade: 19,
		Profissao: "Dev",
	}

	pessoa2 := Pessoa {
		Nome: "Junior",
		Idade: 16,
		Hobby: "Futebol",
	}

	pessoaDescricao := Pessoa {
		Nome: "Antenor",
		Idade: 21,
		Profissao: "Advogado",
		Hobby: "Muay Thai",
		Altura: 1.65,
		Comida: "Carne Assada",
	}
	

	println(pessoa.ListaNomeEIdade())
	println(pessoa2.ListaNomeEIdade2())
	println(pessoaDescricao.DescricaoDaPessoa())
}