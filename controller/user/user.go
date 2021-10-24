package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type UserController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) UserController {
	return UserController{
		usecase: usecase,
	}
}

// swagger:operation GET /me GetUserBySub
// Get user by sub.
// Return user
// ---
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/UserDTO"
//       400:
//         description: Bad Request
func (controller UserController) GetUserBySub(c *gin.Context) {

	sub := c.GetString("sub")

	user, error := controller.usecase.UserUsecase.GetUserBySub(sub)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) AddUser(c *gin.Context) {

	var user contract.UserDTO

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	userDTO, error := controller.usecase.UserUsecase.AddUser(user)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}

// swagger:operation PUT /me ModifyUser
// Put user.
// Return user
// ---
//     Parameters:
//       - name: user
//         in: body
//         schema:
//            $ref: "#/definitions/UserDTO"
//         description: user
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/UserDTO"
//       400:
//         description: Bad Request
func (controller UserController) ModifyUser(c *gin.Context) {

	sub := c.GetString("sub")

	userBySub, err := controller.usecase.UserUsecase.GetUserBySub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	var user contract.UserDTO

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Id = userBySub.Id
	user.ExternalId = userBySub.ExternalId

	userDTO, error := controller.usecase.UserUsecase.ModifyUser(user)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}
