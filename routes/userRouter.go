package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func UserRoutes(incomingRequests *gin.Engine) {
	incomingRequests.GET("/users", controllers.GetUsers())
	incomingRequests.GET("/users/:user_id", controllers.GetUser())
	incomingRequests.POST("/users/login", controllers.Login())
	incomingRequests.POST("/users/signup", controllers.Signup())
}
