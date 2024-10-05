package main

import (
	"api_go_rest/routes"
	"log"
)

func main() {
	log.Printf("Start server")

	routes.HandleRequest()
}
