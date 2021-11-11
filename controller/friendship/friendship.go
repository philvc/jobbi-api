package friendship_controller

import (

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/usecase"
)

type FrienshipControllerInterface interface {
	GetUserFriendships(c *gin.Context)
}

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
// func (controller FriendshipController) GetFriendshipsBySearchId(c *gin.Context) {
// 	searchId := c.Params.ByName("searchId")

// 	status := c.Query("status")

// 	parseStatus, _ := strconv.ParseInt(status, 10, 32)

// 	uintStatus := uint(parseStatus)

// 	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySearchId(searchId, uintStatus)

// 	if error != nil {
// 		c.IndentedJSON(http.StatusBadRequest, error)
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, Friendships)
// }



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

 
}

func (controller FriendshipController) ModifyFriendship(c *gin.Context) {

}

