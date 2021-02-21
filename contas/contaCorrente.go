package contas

import (
	"curso-go-orientacao-a-objetos/clientes"
)

//se as letras iniciais dos campos estiverem minusculas, a visibilidade dos mesmos so serao para esse arquivo
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

	return "Valor do deposito menor que zero", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *ContaCorrente) bool {
	podeTransferir := valorDaTransferencia <= c.saldo && valorDaTransferencia > 0
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
