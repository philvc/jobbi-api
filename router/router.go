package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"

	me_router "github.com/philvc/jobbi-api/router/me"
	organisation_router "github.com/philvc/jobbi-api/router/organisation"
	user_router "github.com/philvc/jobbi-api/router/user"
	search_router "github.com/philvc/jobbi-api/router/search"
)

type Router struct {
	engine     *gin.Engine
	controller controller.Controller
}

func Default(engine *gin.Engine, controller controller.Controller) Router {
	return Router{
		engine:     engine,
		controller: controller,
	}
}

// Initiliases the router with the entire sub-tree
func (router Router) Initiliase() {

	// CORS
	router.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Creates all the routers
	meRouter := me_router.Default(router.controller)
	userRouter := user_router.Default(router.controller)
	organisationRouter := organisation_router.Default(router.controller)
	searchRouter := search_router.Default(router.controller)

	// Creates the api-group
	api := router.engine.Group("")

	// Initialises all the routers
	meRouter.Initialise(api)
	userRouter.Initialise(api)
	organisationRouter.Initialise(api)
	searchRouter.Initialise(api)

}

// Runs the engine
func (router Router) Run() {

	// Runs the engine
	router.engine.Run()
}
