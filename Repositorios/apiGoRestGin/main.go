package main

import (
	"go-rest-api-gin/database"
	"go-rest-api-gin/routes"
)

func main() {

	// models.Alunos = []models.Aluno{
	// 	{Nome: "Patrick", CPF: "00000000000", RG: "470000000"},
	// 	{Nome: "Lucas", CPF: "00000000001", RG: "470000001"},
	// 	{Nome: "Arthur", CPF: "00000000002", RG: "470000002"},
	// }

	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
