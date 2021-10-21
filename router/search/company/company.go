package search_company_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
)

// Name of the endpoint
const endpoint string = "companies"

// Name of the parameter
const parameter string = ":companyId"

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
	child.GET("", middleware.Authorize(routerGroup.controller.CompanyController.GetCompaniesBySearchId))
	child.POST("", middleware.Authorize(routerGroup.controller.CompanyController.AddCompany))

	childParam := child.Group(parameter)
	childParam.GET("", middleware.Authorize(routerGroup.controller.CompanyController.GetCompanyById))
	childParam.PUT("", middleware.Authorize(routerGroup.controller.CompanyController.ModifyCompany))
	childParam.DELETE("", middleware.Authorize(routerGroup.controller.CompanyController.DeleteCompany))
}
