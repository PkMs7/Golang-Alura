package main

import (
	"go-rest-api/db"
	"go-rest-api/routes"
)

func main() {

	// Código antes do ORM
	// models.Personalidades = []models.Personalidade{

	// 	{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
	// 	{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	// }

	// Conexão com o banco de dados via GORM
	db.ConectaComBancoDeDados()

	// Subindo o Servidor
	routes.HandleRequest()

}
