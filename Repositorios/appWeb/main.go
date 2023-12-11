package main

import (
	"appWeb/routes"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
