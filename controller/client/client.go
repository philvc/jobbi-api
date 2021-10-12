package client_controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/usecase"
)

type ClientController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) ClientController {
	return ClientController{
		usecase: usecase,
	}
}

func (controller ClientController) GetClients(c *gin.Context) {

	organisationId := c.Params.ByName("organisationId")

	clients, error := controller.usecase.ClientUsecase.GetClientsByOrganisationId(organisationId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, clients)
}

func (controller ClientController) GetClientById(c *gin.Context) {
	clientId := c.Params.ByName("clientId")
	parsedClientId, _ := strconv.ParseInt(clientId, 10, 32)

	client, error := controller.usecase.ClientUsecase.GetClientById(parsedClientId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}

func (controller ClientController) AddClient(c *gin.Context) {
	var client contract.ClientDTO

	organisationId := c.Params.ByName("organisationId")
	parsedOrganisationId, _ := strconv.ParseInt(organisationId, 10, 32)
	log.Default().Println(organisationId)

	client.OrganisationId = uint(parsedOrganisationId)

	if err := c.BindJSON(&client); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	clientDTO, error := controller.usecase.ClientUsecase.AddClient(client)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientDTO)
}

func (controller ClientController) ModifyClient(c *gin.Context) {
	var client contract.ClientDTO

	if err := c.BindJSON(&client); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	clientDTO, error := controller.usecase.ClientUsecase.ModifyClient(client)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientDTO)
}
