package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // compoe o outro struct, adiciona junto com o atual
	Sallary  Salario
}

// aplica a interface a todos os tipos que possuem Desativar()
// aqui só podemos ter métodos, não podemos ter variaveis.
type Pessoa interface {
	Desativar()
}

type Salario struct {
	Valor int
	Cargo string
}

// Funciona como se fosse um método do Struct. Vira uma função chamada pelo objeto struct.
func (s Salario) Dobrar() {
	s.Valor = s.Valor * 2
	fmt.Printf("O Salario dobrou! Agora é %v", s.Valor)
}

func (c Cliente) Desativar() {

}

func main() {
	marcos := Cliente{
		Nome:  "Marcao",
		Idade: 18,
		Ativo: false,
	}

	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", marcos.Nome, marcos.Idade, marcos.Ativo)
	marcos.Ativo = true
	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", marcos.Nome, marcos.Idade, marcos.Ativo)

	marcos.Cidade = "São Paulo"
	marcos.Endereco.Cidade = "SFO"
	marcos.Sallary.Valor = 1000

	marcos.Sallary.Dobrar()
}
