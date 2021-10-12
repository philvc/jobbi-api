package user_organisation_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/usecase"
)

type UserOrganisationController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) UserOrganisationController {
	return UserOrganisationController{
		usecase: usecase,
	}
}

func (controller UserOrganisationController) AddUserToOrganisation(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Params.ByName("userId"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	var organisation contract.OrganisationDTO

	if err := c.BindJSON(&organisation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	userDTO, error := controller.usecase.UserUsecase.AddUserToOrganisation(uint(userId), organisation)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}
