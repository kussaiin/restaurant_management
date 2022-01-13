package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func TableRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/tables", controllers.GetTables())
	incomingRequest.GET("/tables/:table_id", controllers.GetTable())
	incomingRequest.POST("/tables", controllers.CreateTable())
	incomingRequest.PATCH("/tables/:table_id", controllers.UpdateTable())
}
