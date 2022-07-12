package main

import (
	"github.com/isaque-vieira2019/api-biblia/database"
	"github.com/isaque-vieira2019/api-biblia/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
