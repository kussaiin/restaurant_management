package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func MenuRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/menus", controllers.GetMenus())
	incomingRequest.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRequest.POST("/menus", controllers.CreateMenu())
	incomingRequest.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
