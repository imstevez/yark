package main

import (
	"controllers"
	"log"
	"nicego"
	"opctx"
	"utils"
	"yark/configs"
	"yark/controllers"
	"yark/models"
)

func applicationInit() {
	log.Println("application init...")
	// Create sql DB conn pool
	sqlDB = models.NewSqlDB(configs.SqlDB)
	// Set global context
	ctx = opctx.SetKeyValue(context.Backgroud(), "base_infos", configs.BaseInfos)
	// Get global route
	route = nicego.NewRoute(contx)
	// register routers
	routersInit()
	log.Println("init complete.")
}
