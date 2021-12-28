package search_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type SearchController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) SearchController {
	return SearchController{
		usecase: usecase,
	}
}

// swagger:operation GET /searches/me searches GetMySearch
// Get my search.
// Return my search
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Get my search
//         schema:
//             "$ref": "#/definitions/MySearchDTO"
//       400:
//         description: Bad Request
func (controller SearchController) GetMySearch(c *gin.Context) {

	sub := c.GetString("sub")

	// Get My Searches - Searches by userId
	search, err := controller.usecase.SearchUsecase.GetMySearch(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, search)

}

// swagger:operation GET /searches/shared searches GetMySharedSearches
// Get my shared searches.
// Return my shared searches
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Get my shared searches
//         schema:
//           type: array
//           items:
//             "$ref": "#/definitions/SharedSearchDTO"
//       400:
//         description: Bad Request
func (controller SearchController) GetMySharedSearches(c *gin.Context) {

	// 

	c.IndentedJSON(http.StatusOK, "get shared searches endpoint")

}

// swagger:operation GET /searches/public searches GetMyFollowedSearches
// Get my followed searches.
// Return my followed searches
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Get my followed searches
//         schema:
//           type: array
//           items:
//             "$ref": "#/definitions/FollowedSearchDTO"
//       400:
//         description: Bad Request
func (controller SearchController) GetMyFollowedSearches(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "get followed searches endpoint")

}

func (controller SearchController) GetFriendsSearches(c *gin.Context) {

	sub := c.GetString("sub")

	// Get My Searches - Searches by userId
	searches, err := controller.usecase.SearchUsecase.GetFriendsSearches(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searches)
}

// swagger:operation GET /searches/{searchId} searches GetSearchById
// type id struct
// Get search by id.
// Return search
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/SearchDTO"
//       400:
//         description: Bad Request

func (controller SearchController) GetSearchById(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	search, err := controller.usecase.SearchUsecase.GetSearchById(searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, search)
}

// swagger:operation POST /searches searches AddSearch
// type id struct
// Create search.
// Return search
// ---
//     Parameters:
//       - name: search
//         in: body
//         schema:
//            $ref: "#/definitions/SearchDTO"
//         description: search
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/SearchDTO"
//       400:
//         description: Bad Request

func (controller SearchController) AddSearch(c *gin.Context) {

	var search contract.SearchDTO

	if err := c.BindJSON(&search); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	sub := c.GetString("sub")
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	search.UserID = userDTO.Id

	searchDTO, err := controller.usecase.SearchUsecase.AddSearch(search)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searchDTO)
}

// swagger:operation PUT /searches/{searchId} searches ModifySearch
// type id struct
// Create search.
// Return search
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: search
//         in: body
//         schema:
//            $ref: "#/definitions/SearchDTO"
//         description: search
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/SearchDTO"
//       400:
//         description: Bad Request

func (controller SearchController) ModifySearch(c *gin.Context) {
	var search contract.SearchDTO

	sub := c.GetString("sub")
	searchId := c.Params.ByName("searchId")

	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	search.Id = searchId
	search.UserID = userDTO.Id

	if err := c.BindJSON(&search); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, error := controller.usecase.SearchUsecase.ModifySearch(search)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searchDTO)
}
