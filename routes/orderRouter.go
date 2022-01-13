package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func OrderRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/orders", controllers.GetOrders())
	incomingRequest.GET("/orders/:order_id", controllers.GetOrder())
	incomingRequest.POST("/orders", controllers.CreateOrder())
	incomingRequest.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
