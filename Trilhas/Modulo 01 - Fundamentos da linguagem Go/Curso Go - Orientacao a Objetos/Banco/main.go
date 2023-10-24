package main

import (
	"fmt"
)

type TitularConta struct {
	Nome, CPF, Profissao string
}

type ContaCorrente struct {
	Titular       TitularConta
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

type ContaPoupanca struct {
	Titular                              TitularConta
	NumeroAgencia, NumeroConta, Operacao int
	Saldo                                float64
}

type verificarConta interface {
	Sacar(valor float64) bool
}

func main() {

	contaPoupancaDoPatrick := ContaPoupanca{}

	contaPoupancaDoPatrick.Depositar(1000)
	contaPoupancaDoPatrick.Sacar(50)

	fmt.Println(contaPoupancaDoPatrick.ObterSaldo())

	// contaDaJulia := ContaCorrente{Titular: TitularConta{Nome: "Julia"}, Saldo: 200}
	// contaDoPatrick := ContaCorrente{Titular: TitularConta{Nome: "Patrick"}, Saldo: 0.00}

	// fmt.Println("Saldo conta Julia: ", contaDaJulia.Saldo)
	// transferencia := contaDaJulia.Transferir(100, &contaDoPatrick)
	// fmt.Println(transferencia)
	// fmt.Println("Saldo atual conta Julia: ", contaDaJulia.Saldo)
	// fmt.Println("Saldo atual conta Patrick: ", contaDoPatrick.Saldo)

}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.Saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.Saldo
}

// O comando c *ContaCorrente indica pra onde a função vai olhar para realizar a operação
func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, float64) {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}

}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (bool, float64) {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}

}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {

	podeDepoistar := valorDoDeposito > 0
	if podeDepoistar {
		c.Saldo += valorDoDeposito
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}

}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, float64) {

	podeDepoistar := valorDoDeposito > 0
	if podeDepoistar {
		c.Saldo += valorDoDeposito
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

// contaDoPatrick := ContaCorrente{"Patrick", 555, 123456, 125.30}

// fmt.Println("O titular da conta é: ", contaDoPatrick.titular)
// fmt.Println("O número da agência é: ", contaDoPatrick.numeroAgencia)
// fmt.Println("O número da conta é: ", contaDoPatrick.numeroConta)
// fmt.Println("O saldo da conta é: ", contaDoPatrick.saldo)

// // Criando conta com ponteiro de memória
// var contaDaCris *ContaCorrente
// // Informar dados com o comando New (conforme outras linguagens orientadas a objeto)
// contaDaCris = new(ContaCorrente)
// contaDaCris.titular = "Cris"

// fmt.Println(contaDaCris)

// fmt.Println("O saldo atual é:", contaDoPatrick.saldo)
// status1, valor1 := contaDoPatrick.sacar(150)
// status2, valor2 := contaDoPatrick.depositar(100)
// fmt.Println("O saque foi realizado?", status1)
// fmt.Println("O saldo atual é:", valor1)
// fmt.Println("O depósito foi realizado?", status2)
// fmt.Println("O saldo atual é:", valor2)
