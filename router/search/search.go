package search_router

import (
	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/controller"
	"github.com/philvc/jobbi-api/middleware"
	search_company_router "github.com/philvc/jobbi-api/router/search/company"
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

	// Get my search
	child.GET("/me", middleware.Authorize(routerGroup.controller.SearchController.GetMySearch))

	// Get my shared searches
	child.GET("/shared", middleware.Authorize(routerGroup.controller.SearchController.GetMySharedSearches))

	// Get my followed searches
	child.GET("/public", middleware.Authorize(routerGroup.controller.SearchController.GetMyFollowedSearches))
	
	// Get public 
	child.GET("/explore", middleware.Authorize(routerGroup.controller.SearchController.GetPublicSearches))
	
	
	// Post search
	child.POST("", middleware.Authorize(routerGroup.controller.SearchController.AddSearch))
	
	// With searchId
	childParam := child.Group(parameter)
	
	// Get requester role for search 
	childParam.GET("/role", middleware.Authorize(routerGroup.controller.SearchController.GetSearchRole))
	
	// Update search
	childParam.PUT("", middleware.Authorize(routerGroup.controller.SearchController.ModifySearch))

	// Get search by id
	childParam.GET("", middleware.Authorize(routerGroup.controller.SearchController.GetSearchById))

	// Get search for an invitation
	childParam.GET("/invitations", middleware.Authorize(routerGroup.controller.SearchController.GetSearchByIdForInvitation))

	/* FRIENDS */
	childParam.GET("/friends", middleware.Authorize(routerGroup.controller.SearchController.GetSearchFriends))

	/* PARTICIPANTS */

	// Get search participants
	childParam.GET("/participants", middleware.Authorize(routerGroup.controller.SearchController.GetParticipantsBySearchId))

	/* POSTS */

	// Get search posts
	childParam.GET("/posts", middleware.Authorize(routerGroup.controller.SearchController.GetPostsBySearchId))

	// Add post for search
	childParam.POST("/posts", middleware.Authorize(routerGroup.controller.SearchController.AddPostBySearchId))

	// Edit post by id
	childParam.PUT("/posts/:postId", middleware.Authorize(routerGroup.controller.SearchController.UpdatePostById))

	// Delete post by id
	childParam.DELETE("/posts/:postId", middleware.Authorize(routerGroup.controller.SearchController.DeletePostById))

	/* COMMENTS */
	// Get comments for post
	childParam.GET("/posts/:postId/comments", middleware.Authorize(routerGroup.controller.SearchController.GetCommentsForPost))
	
	// Put comment by id
	childParam.GET("/posts/:postId/comments/:commentId", middleware.Authorize(routerGroup.controller.SearchController.UpdateCommentById))
	
	// Post comment
	childParam.POST("/posts/:postId/comments", middleware.Authorize(routerGroup.controller.SearchController.CreateCommentForPost))
	
	// Delete comment
	childParam.DELETE("/posts/:postId/comments/:commentId", middleware.Authorize(routerGroup.controller.SearchController.DeleteCommentById))

	/* FRIENDSHIPS */
	childParam.POST("/invitations", middleware.Authorize(routerGroup.controller.SearchController.UpsertFriendship))
	childParam.DELETE("/friendships/:friendshipId", middleware.Authorize(routerGroup.controller.SearchController.DeleteFriendshipById))

	/* FOLLOWERS */
	childParam.POST("/followers", middleware.Authorize(routerGroup.controller.SearchController.PostFollower))
	childParam.DELETE("/followers/:followerId", middleware.Authorize(routerGroup.controller.SearchController.DeleteFollower))

	offerGroup := search_offer_router.Default(routerGroup.controller)
	offerGroup.Initialise(childParam)

	companyGroup := search_company_router.Default(routerGroup.controller)
	companyGroup.Initialise(childParam)

	networkGroup := search_network_router.Default(routerGroup.controller)
	networkGroup.Initialise(childParam)
}
