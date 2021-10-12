package user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/controller"
	"github.com/nightborn-be/invoice-backend/middleware"
	user_organisation_router "github.com/nightborn-be/invoice-backend/router/user/organisation"
)

// Name of the endpoint
const endpoint string = "users"

// Name of the parameter
const parameter string = ":userId"

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
	child.POST("", middleware.Authorize(routerGroup.controller.UserController.AddUser))

	childParam := child.Group(parameter)
	childParam.PUT("", middleware.Authorize(routerGroup.controller.UserController.ModifyUser))

	organisationGroup := user_organisation_router.Default(routerGroup.controller)
	organisationGroup.Initialise(childParam)

}
