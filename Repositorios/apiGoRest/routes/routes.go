package routes

import (
	"go-rest-api/controllers"
	middlaware "go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	//Instância Gorillamux
	r := mux.NewRouter()

	r.Use(middlaware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/insere-personalidade", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/deleta-personalidade/{id}", controllers.DeletaUmaNovaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/edita-personalidade/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
