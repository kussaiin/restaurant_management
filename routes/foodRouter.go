package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func FoodRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/foods", controllers.GetFoods())
	incomingRequest.GET("/foods/:food_id", controllers.GetFood())
	incomingRequest.POST("/foods", controllers.CreateFood())
	incomingRequest.PATCH("/foods/:food_id", controllers.UpdateFood())
}
