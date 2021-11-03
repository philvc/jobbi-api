package user_friendship_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
)

// Name of the endpoint
const endpoint string = "friendships"

// Name of the parameter
const parameter string = ":friendshipId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.FriendshipController.GetFriendshipsBySub))

	childParam := child.Group(parameter)
	childParam.PUT("", middleware.Authorize(routerGroup.controller.FriendshipController.ModifyFriendship))
}