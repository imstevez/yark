package main

import (
	"context"
	"log"
	"net/http"
	"nicego"
	"yark/configs"
	"yark/controllers"
	"yark/models"
)

// Global variables
var (
	sqlDB *sql.DB
	route *nicego.Route
	ctx   context.Context
)

func main() {
	applicationInit() // Initialize application
	defer exit()      // Exit: recycling

	// Start HTTP server
	log.Println("server runing...")
	err := http.ListenAndServe(configs.Server["port"], route)
	if err != nil {
		log.Fatal(err)
	}
}
