package main

import (
	"yark/controllers"
	"yark/models"
)

func routeInit() {
	// Index
	indexCtrl := &controllers.IndexCtrl{
		&models.IndexModel{
			sqlDB,
		},
	}
	route.From("/").Do(controllers.IndexCtrl.Index)
}
