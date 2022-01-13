package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func OrderItemRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/orderItems", controllers.GetOrderItems())
	incomingRequest.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRequest.GET("orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRequest.POST("/orderItems", controllers.CreateOrderItem())
	incomingRequest.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
