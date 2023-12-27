package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page Modularizado")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade

	db.DB.Find(&p)

	json.NewEncoder(w).Encode(p)

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personalidade
	db.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

	// Código antes do ORM
	// for _, personanlidade := range models.Personalidades {

	// 	if strconv.Itoa(personanlidade.Id) == id {

	// 		json.NewEncoder(w).Encode(personanlidade)

	// 	}

	// }
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	db.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)

}

func DeletaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	db.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(&p)

}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	db.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	db.DB.Save(&p)
	json.NewEncoder(w).Encode(&p)

}
