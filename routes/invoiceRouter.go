package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)

func InvoiceRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/invoices", controllers.GetInvoices())
	incomingRequest.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRequest.POST("/invoices", controllers.CreateInvoice())
	incomingRequest.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
