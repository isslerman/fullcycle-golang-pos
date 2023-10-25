package main

import "fmt"

type Cliente struct {
	Nome  string
	Saldo int
}

func (c Cliente) printSaldo() {
	fmt.Printf("O saldo do cliente %s é de %v \n", c.Nome, c.Saldo)
}

// aqui criamos um novo cliente e retornamos o endereço de memória, ponteiro dele.
func newConta(nome string) *Cliente {
	return &Cliente{
		Nome:  nome,
		Saldo: 0,
	}
}

// aqui alteramos o valor do saldo recebendo o ponteiro do Cliente
func (c *Cliente) novoSaldo(valor int) {
	c.Saldo += valor
}

// aqui alteramos o valor local/cópia e não o valor do endereço de memória, ponteiro. ERRADO
func (c Cliente) simularSaldo(valor int) {
	c.Saldo += valor
}

func main() {
	marcos := newConta("Marcos")
	marcos.printSaldo()
	marcos.novoSaldo(500)
	marcos.printSaldo()
	marcos.simularSaldo(1000)
	marcos.printSaldo()
}
