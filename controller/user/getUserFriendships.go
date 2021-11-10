package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /users/{userId}/friendships users GetUserFriendships
// type id struct
// Get user friendships.
// Return friendship
// ---
//     Parameters:
//       - name: userId
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
//             $ref: "#/definitions/SearchWithOwnerAndFriendsDTO"
//       400:
//         description: Bad Request
func (controller UserController) GetUserFriendships(c *gin.Context){

	// Params
	userId := c.Params.ByName("userId")
	status := c.Query("status")

	parseStatus, _ := strconv.ParseInt(status, 10,32)
	intStatus := uint(parseStatus)

	results, err :=	controller.usecase.UserUsecase.GetUserFriendships(userId, intStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(http.StatusOK, results)
}