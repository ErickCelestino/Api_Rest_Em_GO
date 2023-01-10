package routes

import (
	"log"
	"net/http" //Pacote de internet

	"github.com/ErickCelestino/Api_Rest_Em_GO/controllers"
	"github.com/ErickCelestino/Api_Rest_Em_GO/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux" //Pacote para rotas
)

func HandleResquest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonsalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	//Falando para o servidor que qualquer porta pode ser referenciada
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
