package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/test_golang/routes"
)

func main() {
	routes.HandleRequest()
	fmt.Println("Listening API...")
}
