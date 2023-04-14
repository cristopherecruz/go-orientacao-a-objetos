package main

import (
	"fmt"
	"github.com/goorientacaoaobjetos/contas"
)

func main() {

	/*contaGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

	contaGuilherme := ContaCorrente{titular: "Guilherme", numeroConta: 589, numeroAgencia: 123456,
		saldo: 125.5}

	contaBruna := ContaCorrente{"Bruna", 222, 111222,
		200}

	contaBruna2 := &contaBruna

	fmt.Println(contaGuilherme)
	fmt.Println(&contaBruna == contaBruna2)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.titular = "Cris"
	contaCris.saldo = 500

	fmt.Println(*contaCris)*/

	/*contaSilvia := ContaCorrente{"Silva", 0, 0, 500}
	fmt.Println(contaSilvia.sacar(400))
	fmt.Println(contaSilvia.saldo)

	fmt.Println(contaSilvia.depositar(500))
	fmt.Println(contaSilvia.saldo)*/

	/*contaSilvia := contas.ContaCorrente{Titular: "Silva", Saldo: 300}
	contaGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaSilvia.Transferir(200, &contaGustavo)

	fmt.Println(status)
	fmt.Println(contaSilvia.Saldo)
	fmt.Println(contaGustavo.Saldo)*/

	/*clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.111.123.12", Profissao: "Desenvolvedor Go"}
	contaBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456}

	fmt.Println(contaBruno)
	contaBruno.Depositar(100)
	fmt.Println(contaBruno.ObterSaldo())*/

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	pagarBoleto(&contaDenis, 60)

	contaPati := contas.ContaCorrente{}
	contaPati.Depositar(500)
	pagarBoleto(&contaPati, 400)

	fmt.Println(contaDenis.ObterSaldo())
	fmt.Println(contaPati.ObterSaldo())

}

func pagarBoleto(conta VerificarConta, valor float64) {
	conta.Sacar(valor)
}

type VerificarConta interface {
	Sacar(valor float64) string
}
