package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {

	conexao := "user=postgres dbname=golang_alura password=teste123 host=localhost sslmode=disable"

	dbp, err := sql.Open("postgres", conexao)
	if err != nil {

		panic(err.Error())

	}

	return dbp

}
