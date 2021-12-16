package search_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
	search_company_router "github.com/philvc/jobbi-api/router/search/company"
	search_friendship_router "github.com/philvc/jobbi-api/router/search/friendship"
	search_network_router "github.com/philvc/jobbi-api/router/search/network"
	search_offer_router "github.com/philvc/jobbi-api/router/search/offer"
)

// Name of the endpoint
const endpoint string = "searches"

// Name of the parameter
const parameter string = ":searchId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.SearchController.GetSearches))
	// Get friends searches
	child.GET("/friends", middleware.Authorize(routerGroup.controller.SearchController.GetFriendsSearches))
	child.POST("", middleware.Authorize(routerGroup.controller.SearchController.AddSearch))
	childParam := child.Group(parameter)
	childParam.PUT("", middleware.Authorize(routerGroup.controller.SearchController.ModifySearch))
	childParam.GET("", middleware.Authorize(routerGroup.controller.SearchController.GetSearchById))

	offerGroup := search_offer_router.Default(routerGroup.controller)
	offerGroup.Initialise(childParam)

	companyGroup := search_company_router.Default(routerGroup.controller)
	companyGroup.Initialise(childParam)

	networkGroup := search_network_router.Default(routerGroup.controller)
	networkGroup.Initialise(childParam)

	friendshipGroup := search_friendship_router.Default(routerGroup.controller)
	friendshipGroup.Initialise(childParam)
}
