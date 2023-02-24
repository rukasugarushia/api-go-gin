package main

import (
	"github.com/rukasugarushia/api-go-gin/database"
	"github.com/rukasugarushia/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
