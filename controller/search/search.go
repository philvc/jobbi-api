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

	// Check params
	if sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
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

	// Params
	sub := c.GetString("sub")

	// Check params
	if sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
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

	// Params
	sub := c.GetString("sub")

	// Check params
	if sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
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

	// Get params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
	search, err := controller.usecase.SearchUsecase.GetSearchById(searchId, sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Return response with status 200
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

	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check params
	if searchId == ""  || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
	posts, err := controller.usecase.SearchUsecase.GetPostsBySearchId(sub, searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

// swagger:operation GET /searches/{searchId}/participants searches GetSearchParticipants
// type id struct
// Get participants by search id.
// Return participants
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
//             $ref: "#/definitions/ParticipantDTOForSearchById"
//       400:
//         description: Bad Request

func (controller SearchController) GetParticipantsBySearchId(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check params
	if searchId == ""  || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
	response, err := controller.usecase.SearchUsecase.GetParticipantsBySearchId(sub, searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, response)
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
	sub := c.GetString("sub")

	// Check params
	// Check sub param
	if sub == ""{
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return 
	}

	// Check body params
	if err := c.BindJSON(&search); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Check user
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

	// Call usecase
	searchDTO, err := controller.usecase.SearchUsecase.AddSearch(sub, searchDto)

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
//            $ref: "#/definitions/PutSearchRequestDTO"
//         description: search
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/PutSearchResponseDTO"
//       400:
//         description: Bad Request

func (controller SearchController) ModifySearch(c *gin.Context) {
	var requestSearchDTO contract.PutSearchRequestDTO

	// Params
	sub := c.GetString("sub")
	searchId := c.Params.ByName("searchId")

	// Check params
	if sub == "" || searchId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return 
	}

	// Check user
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Check body params
	if err := c.BindJSON(&requestSearchDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody))
		return
	}

	// Map request dto with usecase dto
	searchDTO := contract.SearchDTO{
		Id:          searchId,
		UserID:      userDTO.Id,
		Title:       requestSearchDTO.Title,
		Sector:      requestSearchDTO.Sector,
		Type:        requestSearchDTO.Type,
		Description: requestSearchDTO.Description,
		Tags:        requestSearchDTO.Tags,
	}

	// Call usecase
	search, err := controller.usecase.SearchUsecase.ModifySearch(searchDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Map usecase dto with response dto
	responseSearchDTO := contract.PutSearchResponseDTO{
		Id:          search.Id,
		UserID:      search.UserID,
		Title:       search.Title,
		Sector:      search.Sector,
		Type:        search.Type,
		Description: search.Description,
		Tags:        search.Tags,
	}

	c.IndentedJSON(http.StatusOK, responseSearchDTO)
}

// swagger:operation POST /searches/{searchId}/posts searches AddPostForSearch
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
//       - name: post
//         in: body
//         schema:
//            $ref: "#/definitions/AddPostRequestDTO"
//         description: post
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/AddPostResponseDTO"
//       400:
//         description: Bad Request
func (controller SearchController) AddPostBySearchId(c *gin.Context) {

	// Params
	sub := c.GetString("sub")
	searchId := c.Params.ByName("searchId")

	var postRequest contract.AddPostRequestDTO

	// Check params
	if sub == "" || searchId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return 
	}

	// Check body params
	if err := c.BindJSON(&postRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody))
		return
	}
	
	// Check user
	userDto, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	
	// Map request dto with post dto
	postDto := contract.PostDTO{
		Id:                 "",
		Title:              postRequest.Title,
		Description:        postRequest.Description,
		Type:               postRequest.Type,
		Url:                postRequest.Url,
		UserID:             userDto.Id,
		SearchID:           searchId,
		Tags:               []string{},
		ContactFirstName:   "",
		ContactLastName:    "",
		ContactEmail:       "",
		CompanyName:        "",
		CompanyEmail:       "",
		CompanyPhoneNumber: 0,
		CompanyAddress:     "",
		CompanyUrl:         "",
		ContactPhoneNumber: 0,
	}

	// Call usecase
	postResponseDTO, err := controller.usecase.SearchUsecase.AddPost(&postDto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, postResponseDTO)
}

// swagger:operation PUT /searches/{searchId}/posts/{postId} searches UpdatePostById
// type id struct
// Update post.
// Return post
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: postId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: post
//         in: body
//         schema:
//            $ref: "#/definitions/UpdatePostRequestDTO"
//         description: post
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/UpdatePostResponseDTO"
//       400:
//         description: Bad Request
func (controller SearchController) UpdatePostById(c *gin.Context) {

	// Params
	sub := c.GetString("sub")
	postId := c.Params.ByName("postId")
	searchId := c.Params.ByName("searchId")
	
	var postRequestDto contract.UpdatePostRequestDTO

	// Check params
	if sub == "" || searchId == "" || postId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return 
	}

	// Check body matches dto
	if err := c.BindJSON(&postRequestDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody).Error())
		return
	}

	// Get user
	userDto, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Map request dto to usecase dto
	postDto := contract.PostDTO{
		Id:                 postId,
		Title:              postRequestDto.Title,
		Description:        postRequestDto.Description,
		Type:               postRequestDto.Type,
		Url:                postRequestDto.Url,
		UserID:             userDto.Id,
		SearchID:           searchId,
		Tags:               []string{},
		ContactFirstName:   "",
		ContactLastName:    "",
		ContactEmail:       "",
		CompanyName:        "",
		CompanyEmail:       "",
		CompanyPhoneNumber: 0,
		CompanyAddress:     "",
		CompanyUrl:         "",
		ContactPhoneNumber: 0,
	}

	// Call usecase
	postResponseDto, err := controller.usecase.SearchUsecase.UpdatePostById(&postDto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, postResponseDto)
}

// swagger:operation DELETE /searches/{searchId}/posts/{postId} searches DeletePostById
// type id struct
// Update post.
// Return post
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: postId
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
//            type:  boolean
//       400:
//         description: Bad Request
func (controller SearchController) DeletePostById(c *gin.Context) {

	// Params
	sub := c.GetString("sub")
	searchId := c.Params.ByName("searchId")
	postId := c.Params.ByName("postId")

	// Check params
	if sub == "" || searchId == "" || postId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return 
	}

	// Call usecase
	ok, err := controller.usecase.SearchUsecase.DeletePostById(sub, searchId, postId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, ok)

}
