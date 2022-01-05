package user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
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


}
