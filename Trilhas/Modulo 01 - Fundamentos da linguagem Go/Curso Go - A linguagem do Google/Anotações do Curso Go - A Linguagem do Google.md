# Tópicos do Curso: Go - A liguagem do Google

## 1.0 Introdução ao Go

- Para instalação basta acessar o instalador no site [Go Dev](https://go.dev/)

**Hello World em Go**

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

## 2.0 Trabalho com variáveis

- Declarando variáveis tipadas

```
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int = 24
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}

```

- Inferindo variárveias

```
package main

import "fmt"
import "reflect"

func main() {
    var nome = "Douglas"
    var idade = 24
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}

```

## 3.0 Controlando o fluxo do script

- Fluxo com IF

```
nome := "Patrick"
	versao := 1.3

	fmt.Println("Olá", nome)
	fmt.Println("A versão desse sistema é a", versao)
	fmt.Println("Para prosseguir, escolha uma opção:")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Sistema")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {

		fmt.Println("Monitorando...")

	} else if comando == 2 {

		fmt.Println("Exibindo logs...")

	} else if comando == 0 {

		fmt.Println("Saindo do sistema...")

	} else {
		fmt.Println("Não conheço este comando")
	}

```

- Fluxo com Switch

```
	nome := "Patrick"
	versao := 1.3

	fmt.Println("Olá", nome)
	fmt.Println("A versão desse sistema é a", versao)
	fmt.Println("Para prosseguir, escolha uma opção:")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Sistema")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo do sistema...")

	default:
		fmt.Println("Não conheço este comando")
	}

```

- Fluxo com funções

```
	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo do sistema...")

	default:
		fmt.Println("Não conheço este comando")
	}

func exibeIntroducao() {
	nome := "Patrick"
	versao := 1.3

	fmt.Println("Olá", nome)
	fmt.Println("A versão desse sistema é a", versao)
}

func exibeMenu() {
	fmt.Println("Para prosseguir, escolha uma opção:")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Sistema")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}


```

## 4.0 Fazendo requisições para a web

- Requisiões web com biblioteca nativa do Go

```
import "net/http"

func inicarMonitoramento() {

	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "está com problema. Status Code: ", res.StatusCode)
	}

}

```

## 5.0 As principais coleções do Go

- Array

```
var frutas [4] string
frutas[0] = "Abacaxi"
frutas[1] = "Laranja"
frutas[2] = "Morango"

```

- Slice

```
var frutas [] string {"Abacaxi", "Laranja", "Morango"}

```

- Percorrendo uma lista (ambos os códigos executam a mesma função)

```
for i := 0; i < len(frutas); i++ {
	fmt.PrintLn(frutas[i])
}

for i, fruta := range frutas {
	fmt.Println("Estou passando na posição", i, "a fruta:", fruta)
}

```

## 6.0 Lendo arquivos

- Lendo arquivos

```
func lerSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open(enderecoSites)
	//arquivo, _ := ioutil.ReadFile(enderecoSites)

	if err != nil {
		fmt.Println("Ocorreu o erro na abertura do arquivo", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	return sites
}

```

## 7.0 Escrevendo arquivos

- 

```


```