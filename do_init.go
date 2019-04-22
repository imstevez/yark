package main

import (
	"context"
	"fmt"
	"log"
	"nicego"
	"yark/configs"
	"yark/models"
)

func init() {
	log.Println("init...")

	// Load configs
	conf = configs.Configs

	// Initialize route
	route = nicego.NewRoute(context.Background())

	// Create databae connections pool
	sqlConf := conf["sql_db"].(map[string]interface{})
	if db, err = models.NewSqlDB(sqlConf); err != nil {
		panic(fmt.Sprintf("application init: %v", err))
	}
}
