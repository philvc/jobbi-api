package search_offer_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
)

// Name of the endpoint
const endpoint string = "offers"

// Name of the parameter
const parameter string = ":offerId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.OfferController.GetOffersBySearchId))
	child.POST("", middleware.Authorize(routerGroup.controller.OfferController.AddOffer))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.OfferController.GetOfferById))
}
