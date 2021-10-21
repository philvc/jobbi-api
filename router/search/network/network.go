package search_network_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
)

// Name of the endpoint
const endpoint string = "networks"

// Name of the parameter
const parameter string = ":networkId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.NetworkController.GetNetworksBySearchId))
	child.POST("", middleware.Authorize(routerGroup.controller.NetworkController.AddNetwork))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.NetworkController.GetNetworkById))
	childParam.PUT("", middleware.Authorize(routerGroup.controller.NetworkController.ModifyNetwork))
	childParam.DELETE("", middleware.Authorize(routerGroup.controller.NetworkController.DeleteNetwork))
}
