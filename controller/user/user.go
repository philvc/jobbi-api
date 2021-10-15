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

func (controller UserController) GetUserBySub(c *gin.Context) {

	userId := c.GetString("sub")

	user, error := controller.usecase.UserUsecase.GetUserBySub(userId)

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

func (controller UserController) ModifyUser(c *gin.Context) {
	var user contract.UserDTO

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	userDTO, error := controller.usecase.UserUsecase.ModifyUser(user)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}
