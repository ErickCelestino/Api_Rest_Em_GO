package routes

import (
	"log"
	"net/http" //Pacote de internet

	"github.com/ErickCelestino/Api_Rest_Em_GO/controllers"
	"github.com/gorilla/mux" //Pacote para rotas
)

func HandleResquest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonsalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
