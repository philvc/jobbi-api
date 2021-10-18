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

// swagger:operation GET /searches searches GetSearches
// Get all searches.
// Return all searches
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: All the searches
//         schema:
//           type: array
//           items:
//             "$ref": "#/definitions/SearchDTO"
//       400:
//         description: Bad Request

func (controller SearchController) GetSearches(c *gin.Context) {

	sub := c.GetString("sub")

	// Get My Searches - Searches by userId
	searches, err := controller.usecase.SearchUsecase.GetSearchesByUserSub(sub)

	// Get Searches where I am invited - Searches by friendshipId
	// searches, err := controller.usecase.SearchUsecase.GetSearchesByFriendshipId(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searches)
}

// swagger:operation GET /searches/{searchId} search GetSearchById
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

// swagger:operation POST /searches search AddSearch
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

func (controller SearchController) ModifySearch(c *gin.Context) {
	var search contract.SearchDTO

	if err := c.BindJSON(&search); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	userDTO, error := controller.usecase.SearchUsecase.ModifySearch(search)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}