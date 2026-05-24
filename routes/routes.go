package routes

import (
	"echo-rest-api-mysql/domain/item/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterItemRoutes(e *echo.Echo, itemController *controllers.ItemController) {
	apiV1 := e.Group("/api/v1")
	items := apiV1.Group("/items")

	items.POST("", itemController.Create)
	items.GET("", itemController.GetAll)
	items.GET("/:id_item", itemController.GetByID)
	items.PUT("/:id_item", itemController.Update)
	items.DELETE("/:id_item", itemController.Delete)
}
