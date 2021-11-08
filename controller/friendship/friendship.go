package friendship_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type FriendshipController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) FriendshipController {
	return FriendshipController{
		usecase: usecase,
	}
}

// swagger:operation GET /searches/{searchId}/friendships friendships GetFriendshipsBySearchId
// type id struct
// Get friendships by searchId.
// Return friendship
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: status
//         in: query
//         type: number
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
//             $ref: "#/definitions/FriendshipDTO"
//       400:
//         description: Bad Request
func (controller FriendshipController) GetFriendshipsBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	status := c.Query("status")

	parseStatus, _ := strconv.ParseInt(status, 10, 32)

	uintStatus := uint(parseStatus)

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySearchId(searchId, uintStatus)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Friendships)
}

// swagger:operation GET /searches/{searchId}/friendships friendships GetFriendshipsByUserId
// type id struct
// Get user's friends quests is friendship status is 1 or get user's friendships requests if friendship is 0.
// Return friendship
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: status
//         in: query
//         type: number
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
//             $ref: "#/definitions/FriendshipDTO"
//       400:
//         description: Bad Request
func (controller FriendshipController) GetFriendshipsBySub(c *gin.Context) {
	sub := c.GetString("sub")
	status := c.Query("status")


	parseStatus, _ := strconv.ParseInt(status, 10, 32)

	uintStatus := uint(parseStatus)

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySub(sub, uintStatus)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Friendships)
}

// swagger:operation POST /searches/{searchId}/friendships friendships AddFriendship
// type id struct
// Create friendship.
// Return friendship
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: friendship
//         in: body
//         schema:
//            $ref: "#/definitions/FriendshipDTO"
//         description: friendship
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/FriendshipDTO"
//       400:
//         description: Bad Request
func (controller FriendshipController) AddFriendship(c *gin.Context) {

	body := c.BindJSON(c)

	var Friendship contract.FriendshipDTO

	if err := c.BindJSON(&Friendship); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, err := controller.usecase.SearchUsecase.GetSearchById(searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	Friendship.SearchId = searchDTO.Id


	FriendshipDTO, err := controller.usecase.FriendshipUsecase.AddFriendship(Friendship)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, FriendshipDTO)
}

func (controller FriendshipController) ModifyFriendship(c *gin.Context) {
	var Friendship contract.FriendshipDTO

	if err := c.BindJSON(&Friendship); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	FriendshipDTO, error := controller.usecase.FriendshipUsecase.ModifyFriendship(Friendship)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, FriendshipDTO)
}
