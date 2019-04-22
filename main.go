package main

import (
	"database/sql"
	"log"
	"net/http"
	"nicego"
)

// Global variables
var (
	conf  map[string]interface{}
	route *nicego.Route
	sqlDB *sql.DB
)

func main() {
	// Exist application
	defer doExit()

	// Start HTTP server
	port := conf["server"].(map[string]interface{})["port"].(string)
	log.Printf("server runing at %s...\n", port)
	err := http.ListenAndServe(prot, route)
	if err != nil {
		log.Fatal(err)
	}
}
