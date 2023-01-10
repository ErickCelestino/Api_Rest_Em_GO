package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ErickCelestino/Api_Rest_Em_GO/database"
	"github.com/ErickCelestino/Api_Rest_Em_GO/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonsalidades(w http.ResponseWriter, r *http.Request) {
	//Criando uma lista de personalidades para colocar as que tem no banco de dados
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//Fazendo um filtro na personalidade pelo id
	var personalidade models.Personalidade
	//Pergando o primeiro registro referente ao id
	database.DB.First(&personalidade, id)
	//Mostrando a personsalidade na tela
	json.NewEncoder(w).Encode(personalidade)
}

func CriaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	//Json decoder fazer um decodificação para criar algo no banco
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	//Salvando no banco
	database.DB.Create(&novaPersonalidade)
	//Mostrando a nova personsalidade na tela
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	//Editando a personsalidade no banco de dados
	json.NewDecoder(r.Body).Decode(&personalidade)
	//Salvando ela no banco e exibindo
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
