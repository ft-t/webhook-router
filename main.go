package main

import (
	"fmt"
	"webhook-router/configuration"
	"webhook-router/router"
)

func main() {
	conf := configuration.GetConfiguration()

	router.InitRouter(&conf)

	fmt.Println(conf)
}
