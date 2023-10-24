# Tópicos do Curso: Go - Orientação a Objetos

## 1.0 Minha primeira struct

```
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	var contaDoPatrick ContaCorrente = ContaCorrente{"Patrick", 555, 123456, 125.30}

	fmt.Println("O titular da conta é: ", contaDoPatrick.titular)
	fmt.Println("O número da agência é: ", contaDoPatrick.numeroAgencia)
	fmt.Println("O número da conta é: ", contaDoPatrick.numeroConta)
	fmt.Println("O saldo da conta é: ", contaDoPatrick.saldo)
}

```

## 2.0 Referência, ponteiro e método

- Referência e ponteiro
```
	// Criando conta com ponteiro de memória
	var contaDaCris *ContaCorrente
	// Informar dados com o comando New (conforme outras linguagens orientadas a objeto)
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	fmt.Println(contaDaCris)

```

- Métodos de saque e depósito

```
func main() {

	contaDoPatrick := ContaCorrente{"Patrick", 555, 123456, 125.30}

	fmt.Println("O saldo atual é:", contaDoPatrick.saldo)
	fmt.Println("O saque foi realizado?", contaDoPatrick.sacar(150))
	fmt.Println("O saldo atual é:", contaDoPatrick.saldo)
	fmt.Println("O saque foi realizado?", contaDoPatrick.depositar(100))
	fmt.Println("O saldo atual é:", contaDoPatrick.saldo)


}

// O comando c *ContaCorrente indica pra onde a função vai olhar para realizar a operação
func (c *ContaCorrente) sacar(valorDoSaque float64) (status bool) {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (status bool) {

	podeDepoistar := valorDoDeposito > 0
	if podeDepoistar {
		c.saldo += valorDoDeposito
		return true
	} else {
		return false
	}

}


```

## 3.0 Retornos, pacotes e visibilidade

```
func main() {

	contaDoPatrick := ContaCorrente{"Patrick", 555, 123456, 125.30}

	fmt.Println("O saldo atual é:", contaDoPatrick.saldo)
	status1, valor1 := contaDoPatrick.sacar(150)
	status2, valor2 := contaDoPatrick.depositar(100)
	fmt.Println("O saque foi realizado?", status1)
	fmt.Println("O saldo atual é:", valor1)
	fmt.Println("O depósito foi realizado?", status2)
	fmt.Println("O saldo atual é:", valor2)

}

// O comando c *ContaCorrente indica pra onde a função vai olhar para realizar a operação
func (c *ContaCorrente) sacar(valorDoSaque float64) (bool, float64) {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return true, c.saldo
	} else {
		return false, c.saldo
	}

}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (bool, float64) {

	podeDepoistar := valorDoDeposito > 0
	if podeDepoistar {
		c.saldo += valorDoDeposito
		return true, c.saldo
	} else {
		return false, c.saldo
	}

}

func (c *ContaCorrente) transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

```

## 4.0 Composição e encapsulamento

```
type TitularConta struct {
	Nome      string
	CPF       string
	Profissao string
}

type ContaCorrente struct {
	Titular       TitularConta
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func main() {

	contaDaJulia := ContaCorrente{Titular: TitularConta{Nome: "Julia"}, Saldo: 200}
	contaDoPatrick := ContaCorrente{Titular: TitularConta{Nome: "Patrick"}, Saldo: 0.00}

	fmt.Println("Saldo conta Julia: ", contaDaJulia.Saldo)
	transferencia := contaDaJulia.Transferir(100, &contaDoPatrick)
	fmt.Println(transferencia)
	fmt.Println("Saldo atual conta Julia: ", contaDaJulia.Saldo)
	fmt.Println("Saldo atual conta Patrick: ", contaDoPatrick.Saldo)

}

```

## 5.0 Interface e novo tipo de conta

```
type ContaPoupanca struct {
	Titular                              TitularConta
	NumeroAgencia, NumeroConta, Operacao int
	Saldo                                float64
}

type verificarConta interface {
	Sacar(valor float64) bool
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

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, float64) {

	podeDepoistar := valorDoDeposito > 0
	if podeDepoistar {
		c.Saldo += valorDoDeposito
		return true, c.Saldo
	} else {
		return false, c.Saldo
	}

}



```