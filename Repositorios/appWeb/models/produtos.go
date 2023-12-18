package models

import (
	"appWeb/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
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

		p.Id = id
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

func DeletaProduto(id string) {

	db := db.ConectaComBancoDeDados()

	deletaDadosNoBanco, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletaDadosNoBanco.Exec(id)

	defer db.Close()

}

func EditaProduto(id string) Produto {

	db := db.ConectaComBancoDeDados()

	dadoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	dadoParaAtualizar := Produto{}

	for dadoDoBanco.Next() {

		var Id, Quantidade int
		var Nome, Descricao string
		var Preco float64

		err = dadoDoBanco.Scan(&Id, &Nome, &Descricao, &Preco, &Quantidade)
		if err != nil {
			panic(err.Error())
		}
		dadoParaAtualizar.Id = Id
		dadoParaAtualizar.Nome = Nome
		dadoParaAtualizar.Descricao = Descricao
		dadoParaAtualizar.Preco = Preco
		dadoParaAtualizar.Quantidade = Quantidade

	}

	defer db.Close()
	return dadoParaAtualizar

}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}

// slice de produtos
// produtos := []Produto{
// 	{Nome: "Camiseta", Descricao: "Logo de RPG", Preco: 24.99, Quantidade: 5},
// 	{"Calção", "Logo combinando", 29.99, 2},
// 	{"Camiseta", "Logo de anime", 24.99, 8},
// }
