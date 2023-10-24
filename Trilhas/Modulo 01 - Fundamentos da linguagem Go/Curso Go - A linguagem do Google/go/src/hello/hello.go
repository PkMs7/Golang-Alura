package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

var enderecoSites = "Trilhas/Módulo 01 - Fundamentos da linguagem Go/Curso Go - A linguagem do Google/go/src/docs/sites.txt"
var enderecoLogs = "Trilhas/Módulo 01 - Fundamentos da linguagem Go/Curso Go - A linguagem do Google/go/src/logs/log.txt"

func main() {

	exibeIntroducao()
	for {

		exibeMenu()
		processaComando()

	}

}

// Programa de Hello World

/*
var nome string = "Patrick"
var versao float32 = 1.2
// Inferência de tipo quando não se declara uma variável
var versao2 = 1.2

// , concatena
fmt.Println("Hello World", nome)
fmt.Println("A versão desse sistema é", versao)
fmt.Println("O tipo da variável versão2 é", reflect.TypeOf(versao2))*/

// ----------------------------------------------------------------------

// Criando funções

/*exibeIntroducao()
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
}*/

/*if comando == 1 {

	fmt.Println("Monitorando...")

} else if comando == 2 {

	fmt.Println("Exibindo logs...")

} else if comando == 0 {

	fmt.Println("Saindo do sistema...")

} else {
	fmt.Println("Não conheço este comando")
}*/

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

func processaComando() {

	switch leComando() {
	case 1:
		inicarMonitoramento()

	case 2:
		imprimeLogs()

	case 0:
		fmt.Println("Saindo do sistema...")
		os.Exit(0)

	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

//------------------------------------------------------------------------------------

// Requisições web

func inicarMonitoramento() {

	fmt.Println("Monitorando...")
	//sites := []string{"https://www.alura.com.br", "https://www.google.com.br", "https://www.youtube.com.br"}

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando o site que está na posição", i, ". O site testado é:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problema. Status Code: ", res.StatusCode)
		registraLog(site, false)
	}

}

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

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile(enderecoLogs, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println("Ocorreu o erro na abertura do arquivo", err)

	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {

	fmt.Println("Exibindo logs...")

	arquivo, err := ioutil.ReadFile(enderecoLogs)

	if err != nil {
		fmt.Println("Ocorreu o erro na abertura do arquivo", err)

	}

	fmt.Println(string(arquivo))

}
