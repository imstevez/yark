package main

import (
	"yark/controllers"
	"yark/models"
)

func init() {
	ctrl := &controllers.IndexCtrl{
		Model: &models.IndexModel{
			DB: sqlDB,
		},
	}

	route.From("/").Do(ctrl.Index)
}
