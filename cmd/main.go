package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/test_golang/config"
	"github.com/Gabriel-Newton-dev/test_golang/packages/routes"
)

func main() {
	config.LoadConfig()
	routes.HandleRequest()
	fmt.Println("Listening API...")
}
