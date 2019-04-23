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
	// Server configs
	serverConf := conf["server"].(map[string]interface{})
	// Run server
	err := http.ListenAndServe(serverConf["port"].(string), route)
	if err != nil {
		log.Fatal(err)
	}
}
