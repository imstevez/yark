package controllers

import (
	"context"
	"fmt"
	"nicego"
	"yark/models"
)

type IndexCtrl struct {
	Model *models.IndexModel
}

func (ctrl *IndexCtrl) Index(ctx context.Context) {
	w, _ := nicego.GetMeta(ctx)
	fmt.Fprintf(w, "Hello, world.\n")
}
