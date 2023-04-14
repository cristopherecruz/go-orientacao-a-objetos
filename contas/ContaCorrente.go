package contas

import "github.com/goorientacaoaobjetos/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {

	isSacar := valor > 0 && valor <= c.saldo
	if isSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {

	isDepositar := valor > 0
	if isDepositar {
		c.saldo += valor
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Deposito não realizado com sucesso, valor menor que zero!", valor
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	isTransferir := valor <= c.saldo && valor > 0
	if isTransferir {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
