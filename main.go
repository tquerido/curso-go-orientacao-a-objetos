package main

import (
	"curso-go-orientacao-a-objetos/clientes"
	"curso-go-orientacao-a-objetos/contas"
	"fmt"
)

func PagarBoleto(conta conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type conta interface {
	Sacar(valor float64) string
}

func main() {
	fmt.Println("Realizando testes no main.go...")

	clientePedro := clientes.Titular{
		Nome:      "Pedro",
		CPF:       "923.456.789-61",
		Profissão: "Arquiteto"}

	contaDoPedro := contas.ContaPoupanca{
		Titular:       clientePedro,
		NumeroAgencia: 123,
		NumeroConta:   987654,
		Operacao:      0,
	}

	contaDoPedro.Depositar(1000)
	PagarBoleto(&contaDoPedro, 200)
	contaDoPedro.Sacar(300)
	fmt.Println("Saldo do Pedro: ", contaDoPedro.ObterSaldo())

	clienteLuiza := clientes.Titular{
		Nome:      "Luiza",
		CPF:       "123.456.789-50",
		Profissão: "Desenvolvedora"}

	contaDaLuiza := contas.ContaCorrente{
		Titular:       clienteLuiza,
		NumeroAgencia: 123,
		NumeroConta:   987654,
	}

	clienteTabata := clientes.Titular{
		Nome:      "Tabata",
		CPF:       "823.456.789-99",
		Profissão: "Engenheira"}

	contaDaTabata := contas.ContaCorrente{
		Titular:       clienteTabata,
		NumeroAgencia: 123,
		NumeroConta:   987654,
	}

	contaDaLuiza.Depositar(1000)
	contaDaLuiza.Sacar(50)
	PagarBoleto(&contaDaLuiza, 50)
	contaDaLuiza.Transferir(100, &contaDaTabata)
	fmt.Println("Saldo da Luiza: ", contaDaLuiza.ObterSaldo())
	fmt.Println("Saldo da Tabata: ", contaDaTabata.ObterSaldo())
}
