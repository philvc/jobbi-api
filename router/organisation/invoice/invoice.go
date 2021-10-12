package organisation_invoice_router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/controller"
	"github.com/nightborn-be/invoice-backend/middleware"
)

// Name of the endpoint
const endpoint string = "invoices"

// Name of the parameter
const parameter string = ":invoiceId"

type RouterGroup struct {
	controller controller.Controller
}

func Default(controller controller.Controller) RouterGroup {
	return RouterGroup{
		controller: controller,
	}
}

// Generates the sub-tree of routes
func (routerGroup RouterGroup) Initialise(parent *gin.RouterGroup) {

	child := parent.Group(endpoint)
	child.GET("", middleware.Authorize(routerGroup.controller.InvoiceController.GetInvoices))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.InvoiceController.GetInvoiceById))
	childParam.GET("/export", middleware.Authorize(routerGroup.controller.ExportController.GetExport))
}
