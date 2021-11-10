package user_controller

import "github.com/gin-gonic/gin"


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
//             $ref: "#/definitions/FriendshipDTO"
//       400:
//         description: Bad Request
func (controller UserController) GetUserFriendships(c *gin.Context){

	

}