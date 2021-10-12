package me_router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/controller"
	"github.com/nightborn-be/invoice-backend/middleware"
)

// Name of the endpoint
const endpoint string = "me"

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

	users := parent.Group(endpoint)
	users.GET("", middleware.Authorize(routerGroup.controller.UserController.GetUserBySub))

}
