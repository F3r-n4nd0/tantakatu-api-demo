package routes

import (
	"api.tantakatu.com/models"
	"api.tantakatu.com/restapi/operations/items"
	"github.com/go-openapi/runtime/middleware"
)

type itemsHandler struct {
}

func NewItemsHandler() *itemsHandler {
	return &itemsHandler{
	}
}

func (handler *itemsHandler) List(params items.ListItemsParams) middleware.Responder {
	return &items.ListItemsOK{
		Payload: []*models.ItemList{},
	}
}