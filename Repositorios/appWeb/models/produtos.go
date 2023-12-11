package models

import (
	"appWeb/db"
)

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos;")
	if err != nil {

		panic(err.Error())

	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()

	return produtos

}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

// slice de produtos
// produtos := []Produto{
// 	{Nome: "Camiseta", Descricao: "Logo de RPG", Preco: 24.99, Quantidade: 5},
// 	{"Calção", "Logo combinando", 29.99, 2},
// 	{"Camiseta", "Logo de anime", 24.99, 8},
// }