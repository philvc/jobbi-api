package organisation_controller

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/usecase"
)

type OrganisationController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) OrganisationController {
	return OrganisationController{
		usecase: usecase,
	}
}

func (controller OrganisationController) GetOrganisations(c *gin.Context) {

	sub := c.GetString("sub")

	organisations, err := controller.usecase.OrganisationUsecase.GetOrganisationsByUserSub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, organisations)
}

func (controller OrganisationController) GetOrganisationById(c *gin.Context) {
	organisationId := c.Params.ByName("organisationId")

	organisation, error := controller.usecase.OrganisationUsecase.GetOrganisationById(organisationId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, organisation)
}

func (controller OrganisationController) AddOrganisation(c *gin.Context) {

	var organisation contract.OrganisationDTO

	if err := c.BindJSON(&organisation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	rmv_space := regexp.MustCompile(`\s+`)
	slug := rmv_space.ReplaceAllString(organisation.Name, "-")
	slug = strings.ToLower(slug)
	organisation.Slug = slug

	sub := c.GetString("sub")
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	organisation.Owner = userDTO.Id

	organisationDTO, err := controller.usecase.OrganisationUsecase.AddOrganisation(organisation)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, organisationDTO)
}

func (controller OrganisationController) ModifyOrganisation(c *gin.Context) {
	var organisation contract.OrganisationDTO

	if err := c.BindJSON(&organisation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	rmv_space := regexp.MustCompile(`\s+`)
	slug := rmv_space.ReplaceAllString(organisation.Name, "-")
	slug = strings.ToLower(slug)
	organisation.Slug = slug

	userDTO, error := controller.usecase.OrganisationUsecase.ModifyOrganisation(organisation)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, userDTO)
}
