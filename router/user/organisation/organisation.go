package user_organisation_router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/controller"
	"github.com/nightborn-be/invoice-backend/middleware"
)

// Name of the endpoint
const endpoint string = "organisation"

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
	child.POST("", middleware.Authorize(routerGroup.controller.UserOrganisationController.AddUserToOrganisation))

}
