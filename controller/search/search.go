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
	if searchId == "" || sub == "" {
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

// swagger:operation GET /searches/{searchId}/posts/{postId} posts GetPostById
// type id struct
// Get post by  id.
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
//            $ref: "#/definitions/PostDTOBySearchId"
//       400:
//         description: Bad Request
func (controller SearchController) GetPostById(c *gin.Context) {

	// Params
	postId := c.Params.ByName("postId")
	sub := c.GetString("sub")

	// Check params
	if postId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams))
		return
	}

	// Call usecase
	post, err := controller.usecase.SearchUsecase.GetPostById(sub, postId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, post)

}

// swagger:operation GET /searches/{searchId}/friends searches GetSearchFriends
// type id struct
// Get friends by search id.
// Return friends
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
//             $ref: "#/definitions/UserDTO"
//       400:
//         description: Bad Request
func (controller *SearchController) GetSearchFriends(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Verify Params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	friends, err := controller.usecase.SearchUsecase.GetFriendsBySearchId(sub, searchId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, friends)
}

// swagger:operation GET /searches/{searchId}/participants searches GetSearchParticipants
// type id struct
// Get participants (friends & followers) by search id.
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
	if searchId == "" || sub == "" {
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
	if sub == "" {
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

// swagger:operation GET /searches/{searchId}/invitations searches GetSearchByIdForInvitation
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
func (controller SearchController) GetSearchByIdForInvitation(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	search, err := controller.usecase.SearchUsecase.GetSearchByIdForInvitation(sub, searchId)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, search)

}

// swagger:operation POST /searches/{searchId}/invitations searches UpsertFriendship
// type id struct
// Upsert friendship.
// Return friendship.
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
//            $ref: "#/definitions/UpsertFriendshipRequestDTO"
//         description: post
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/UpsertFriendshipResponseDTO"
//       400:
//         description: Bad Request
func (controller SearchController) UpsertFriendship(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")
	var requestDto contract.UpsertFriendshipRequestDTO

	// Check Params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	if err := c.BindJSON(&requestDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody).Error())
		return
	}

	// Get User
	userDto, err := controller.usecase.UserUsecase.GetUserBySub(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Map usecase dto from request dto
	friendshipDto := contract.FriendshipDTO{
		Id:       "",
		Type:     requestDto.Type,
		State:    requestDto.State,
		SearchId: searchId,
		UserId:   userDto.Id,
	}

	// Call usecase
	response, err := controller.usecase.SearchUsecase.UpsertFriendship(&friendshipDto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// map controller response from dto
	friendshipResponseDto := contract.UpsertFriendshipResponseDTO{
		Id:       response.Id,
		Type:     response.Type,
		State:    response.State,
		SearchId: response.SearchId,
		UserId:   response.UserId,
	}

	c.IndentedJSON(http.StatusOK, friendshipResponseDto)
}

// swagger:operation DELETE /searches/{searchId}/friendships/{friendshipId} searches DeleteFriendshipById
// type id struct
// Delete friendship.
// Return boolean
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: friendshipId
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
func (controller SearchController) DeleteFriendshipById(c *gin.Context) {
	// Params
	searchId := c.Params.ByName("searchId")
	friendshipId := c.Params.ByName("friendshipId")
	sub := c.GetString("sub")

	// Check params
	if searchId == "" || friendshipId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	ok, err := controller.usecase.SearchUsecase.DeleteFriendshipById(sub, searchId, friendshipId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ok)
}

// swagger:operation POST /searches/{searchId}/followers searches PostFollower
// type id struct
// Post follower.
// Return follower.
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
//            $ref: "#/definitions/FollowerDTO"
//       400:
//         description: Bad Request
func (controller SearchController) PostFollower(c *gin.Context) {
	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check Params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	follower, err := controller.usecase.SearchUsecase.PostFollower(sub, searchId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, follower)

}

// swagger:operation DELETE /searches/{searchId}/followers/{followerId} searches DeleteFollower
// type id struct
// Delete follower.
// Return boolean
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: followerId
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
func (controller SearchController) DeleteFollower(c *gin.Context) {
	// Params
	searchId := c.Params.ByName("searchId")
	followerId := c.Params.ByName("followerId")
	sub := c.GetString("sub")

	// Check Params
	if searchId == "" || followerId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	ok, err := controller.usecase.SearchUsecase.DeleteFollowerById(sub, searchId, followerId)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ok)
}

// swagger:operation GET /searches/explore searches GetPublicSearches
// type id struct
// Get public searches
// Return searches
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//           type: array
//           items:
//             $ref: "#/definitions/PublicSearchDto"
//       400:
//         description: Bad Request
func (controller SearchController) GetPublicSearches(c *gin.Context) {

	// Params
	sub := c.GetString("sub")

	// Check params
	if sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// call usecase
	searches, err := controller.usecase.SearchUsecase.GetPublicSearches(sub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, searches)

}

// swagger:operation GET /searches/{searchId}/role searches GetSearchRole
// type id struct
// Get role by search id.
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
//            type: string
//            enum: [owner, friend, follower, visitor]
//       400:
//         description: Bad Request
func (controller SearchController) GetSearchRole(c *gin.Context) {
	// Params
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	// Check params
	if searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	role, err := controller.usecase.SearchUsecase.GetSearchRole(sub, searchId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, role)
}

// swagger:operation GET /searches/{searchId}/posts/{postId}/comments comments GetCommentsForPost
// type id struct
// Get comments for post
// Return comments
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
//           type: array
//           items:
//             $ref: "#/definitions/CommentForPostDto"
//       400:
//         description: Bad Request
func (controller SearchController) GetCommentsForPost(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	postId := c.Params.ByName("postId")
	sub := c.GetString("sub")

	// Check params
	if searchId == "" || postId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call usecase
	response, err := controller.usecase.SearchUsecase.GetAllCommentsForPost(sub, searchId, postId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusAccepted, response)

}

// swagger:operation PUT /searches/{searchId}/posts/{postId}/comments/{commentId} comments UpdateCommentById
// type id struct
// Update comment for post
// Return comment
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
//       - name: commentId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: post
//         in: body
//         schema:
//            $ref: "#/definitions/CommentUpdateRequestDto"
//         description: post
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//             $ref: "#/definitions/CommentUpdateDto"
//       400:
//         description: Bad Request
func (controller SearchController) UpdateCommentById(c *gin.Context) {
	// Params
	searchId := c.Params.ByName("searchId")
	postId := c.Params.ByName("postId")
	commentId := c.Params.ByName("commentId")
	sub := c.GetString("sub")
	var requestDto contract.CommentUpdateRequestDto

	// Check params
	if searchId == "" || postId == "" || sub == "" || commentId == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	if err := c.BindJSON(&requestDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody).Error())
		return
	}

	// Create request dto
	commentDto := contract.CommentDTO{
		SearchId: searchId,
		PostId:   postId,
		Content:  requestDto.Content,
		Id:       commentId,
	}
	// Call usecase
	response, err := controller.usecase.SearchUsecase.UpdateCommentForPost(sub, &commentDto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, response)

}

// swagger:operation POST /searches/{searchId}/posts/{postId}/comments comments CreateCommentForPost
// type id struct
// Add comment for post
// Return comment
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
//            $ref: "#/definitions/CommentCreateRequestDto"
//         description: post
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//             $ref: "#/definitions/CreateCommentResponseDto"
//       400:
//         description: Bad Request
func (controller SearchController) CreateCommentForPost(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	postId := c.Params.ByName("postId")
	sub := c.GetString("sub")
	var requestDto contract.CommentDTO

	// Check params
	if searchId == "" || postId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	if err := c.BindJSON(&requestDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongBody).Error())
		return
	}

	// Build usecase dto
	requestDto.PostId = postId
	requestDto.SearchId = searchId

	// Call usecase
	response, err := controller.usecase.SearchUsecase.CreateCommentForPost(sub, &requestDto)
	if err != nil {

		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, response)

}

// swagger:operation DELETE /searches/{searchId}/posts/{postId}/comments/{commentId} comments DeteCommentById
// type id struct
// Delete comment by id
// Return boolean
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
//       - name: commentId
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
func (controller SearchController) DeleteCommentForPostById(c *gin.Context) {

	// Params
	commentId := c.Params.ByName("commentId")
	postId := c.Params.ByName("postId")
	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	if commentId == "" || postId == "" || searchId == "" || sub == "" {
		c.IndentedJSON(http.StatusBadRequest, errors.New(constant.ErrorWrongParams).Error())
		return
	}

	// Call repo
	ok, err := controller.usecase.SearchUsecase.DeleteCommentForPost(sub, searchId, postId, commentId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ok)

}
