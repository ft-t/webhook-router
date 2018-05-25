package main

import (
	"fmt"
	"webhook-router/configuration"
)

func main() {
	conf := configuration.GetConfiguration()

	fmt.Println(conf)
}
