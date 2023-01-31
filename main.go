package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Jeferson Santos Oliveira", Historia: "lorem Ipsum 1"},
		{Id: 2, Nome: "José Oliveira", Historia: "lorem Ipsum 2"},
	}

	fmt.Println("Iniciando o servidor com Go")
	routes.HandleRequest()
}
