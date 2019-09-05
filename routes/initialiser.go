package routes

import (
	"api.tantakatu.com/restapi/operations"
	"api.tantakatu.com/restapi/operations/items"
)

func BindRoutes(api *operations.TantakatuAPI) {

	itemsHandler := NewItemsHandler()
	api.ItemsListItemsHandler = items.ListItemsHandlerFunc(itemsHandler.List)

}