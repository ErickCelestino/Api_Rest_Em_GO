package main

import (
	"fmt" //Pacote que mostra as coisas na tela

	"github.com/ErickCelestino/Api_Rest_Em_GO/database"
	"github.com/ErickCelestino/Api_Rest_Em_GO/models"
	"github.com/ErickCelestino/Api_Rest_Em_GO/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
