package main

import (
	"fmt"
	"golang-microservice-template/config"
	"golang-microservice-template/internal/app"
)

func main() {
	fmt.Println("Starting sendpush-service")

	appConfiguration, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(appConfiguration)
	}

	// Run
	app.Run(&appConfiguration)
}
