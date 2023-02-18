package main

import (
	"fmt"

	"github.com/leandro-ikehara/Estudos-GoLang/go-rest-api/database"
	"github.com/leandro-ikehara/Estudos-GoLang/go-rest-api/model"
	"github.com/leandro-ikehara/Estudos-GoLang/go-rest-api/routes"
)

func main() {
	model.Personalidades = []model.Personalidade{
		{Id: 1, Nome: "Washington Luis", Historia: "Historia 1"},
		{Id: 2, Nome: "Faria Lima", Historia: "Historia 2"},
	}
	database.ConectaBancoDeDados()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
