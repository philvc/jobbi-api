package organisation_client_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
	organisation_client_invoice_router "github.com/philvc/jobbi-api/router/organisation/client/invoice"
)

// Name of the endpoint
const endpoint string = "clients"

// Name of the parameter
const parameter string = ":clientId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.ClientController.GetClients))
	child.POST("", middleware.Authorize(routerGroup.controller.ClientController.AddClient))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.ClientController.GetClientById))
	childParam.PUT("", middleware.Authorize(routerGroup.controller.ClientController.ModifyClient))

	invoiceGroup := organisation_client_invoice_router.Default(routerGroup.controller)
	invoiceGroup.Initialise(childParam)
}
