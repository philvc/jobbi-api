package organisation_router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/controller"
	"github.com/nightborn-be/invoice-backend/middleware"
	organisation_client_router "github.com/nightborn-be/invoice-backend/router/organisation/client"
	organisation_invoice_router "github.com/nightborn-be/invoice-backend/router/organisation/invoice"
)

// Name of the endpoint
const endpoint string = "organisations"

// Name of the parameter
const parameter string = ":organisationId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.OrganisationController.GetOrganisations))
	child.POST("", middleware.Authorize(routerGroup.controller.OrganisationController.AddOrganisation))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.OrganisationController.GetOrganisationById))
	childParam.PUT("", middleware.Authorize(routerGroup.controller.OrganisationController.ModifyOrganisation))

	clientGroup := organisation_client_router.Default(routerGroup.controller)
	clientGroup.Initialise(childParam)

	invoiceGroup := organisation_invoice_router.Default(routerGroup.controller)
	invoiceGroup.Initialise(childParam)
}
