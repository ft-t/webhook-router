package main

import (
	"webhook-router/configuration"
	"webhook-router/db"
	"webhook-router/router"
)

func main() {
	conf := configuration.GetConfiguration()

	db.InitDb(&conf)
	router.InitRouter(&conf)
}
