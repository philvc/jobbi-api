package search_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	constant "github.com/philvc/jobbi-api/constants"
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

	sub := c.GetString("sub")

	// Get Shared searches
	sharedSearches, err := controller.usecase.SearchUsecase.GetSharedSearches(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, sharedSearches)

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

	sub := c.GetString("sub")

	// Get Shared searches
	followedSearches, err := controller.usecase.SearchUsecase.GetFollowedSearches(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, followedSearches)

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
//            $ref: "#/definitions/SearchDTOById"
//       400:
//         description: Bad Request

func (controller SearchController) GetSearchById(c *gin.Context) {
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check search access rights
	ok := controller.hasSearchAccess(sub, searchId)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorMissingAccess).Error())
		return
	}

	search, err := controller.usecase.SearchUsecase.GetSearchById(searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, search)
}

// swagger:operation GET /searches/{searchId}/posts searches GetSearchPosts
// type id struct
// Get posts by search id.
// Return posts
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
//           type: array
//           items:
//             $ref: "#/definitions/PostDTOBySearchId"
//       400:
//         description: Bad Request

func (controller SearchController) GetPostsBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check search access rights
	ok := controller.hasSearchAccess(sub, searchId)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorMissingAccess).Error())
		return
	}

	search, err := controller.usecase.SearchUsecase.GetPostsBySearchId(searchId)

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
//            $ref: "#/definitions/PostSearchRequestDTO"
//         description: search
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/PostSearchResponseDTO"
//       400:
//         description: Bad Request

func (controller SearchController) AddSearch(c *gin.Context) {

	var search contract.PostSearchRequestDTO

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

	// map post request dto with usecase dto
	searchDto := contract.SearchDTO{
		Id:          "",
		Description: search.Description,
		Title:       search.Title,
		UserID:      userDTO.Id,
		Tags:        search.Tags,
		Sector:      search.Sector,
		Type:        search.Type,
	}

	// Add Search usecase
	searchDTO, err := controller.usecase.SearchUsecase.AddSearch(searchDto)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// map post response dto with usecase dto
	postResponseDto := contract.PostSearchResponseDTO{
		Id:          searchDTO.Id,
		UserID:      searchDTO.UserID,
		Title:       searchDTO.Title,
		Description: search.Description,
		Sector:      searchDTO.Sector,
		Tags:        searchDTO.Tags,
		Type:        searchDTO.Type,
	}

	c.IndentedJSON(http.StatusOK, postResponseDto)
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

func (controller SearchController) hasSearchAccess(sub string, searchId string) bool {

	// Check if search is public
	isPublic := controller.usecase.SearchUsecase.IsPublic(searchId)
	if !isPublic {

		// Check if user is owner
		isOwner := controller.usecase.SearchUsecase.IsOwner(sub, searchId)
		if !isOwner {

			// Check if user is friend or follower
			isFriend := controller.usecase.SearchUsecase.IsFriend(sub, searchId)

			if !isFriend {
				return false
			}
		}
	}
	return true
}
