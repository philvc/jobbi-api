package friendship_controller

import (
	"log"
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

	log.Print("params", c)

	parseStatus, _ := strconv.ParseInt("0", 10,32)

	uintStatus := uint(parseStatus)

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySearchId(searchId, uintStatus)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Friendships)
}
func (controller FriendshipController) GetFriendshipsBySub(c *gin.Context) {
	sub := c.GetString("sub")

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySub(sub)

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

	searchId := c.Params.ByName("searchId")

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

	// @todo connect with iam and create account
	// Get User by email
	user, _ := controller.usecase.UserUsecase.GetUserByEmail(Friendship.Email)

	if user != nil {
		Friendship.UserId = user.Id
	}

	// // If no user, create new account supabase & crate new user
	// if err != nil {

	// 	// create temporary account in iam

	// 	// map friendship to user
	// 	var newUser contract.UserDTO

	// 	newUser.Email = Friendship.Email
	// 	newUser.FirstName = Friendship.FirstName
	// 	newUser.LastName = Friendship.LastName

	// 	new, err := controller.usecase.UserUsecase.AddUser(newUser)

	// 	if err != nil {

	// 		c.IndentedJSON(http.StatusBadRequest, err.Error())
	// 		return
	// 	}

	// 	Friendship.UserId = new.Id
	// }

	// save userId in Friendship.UserId

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
