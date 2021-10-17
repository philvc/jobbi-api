package offer_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type OfferController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) OfferController {
	return OfferController{
		usecase: usecase,
	}
}

func (controller OfferController) GetOffersBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	offers, error := controller.usecase.OfferUsecase.GetOffersBySearchId(searchId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, offers)
}

func (controller OfferController) GetOfferById(c *gin.Context) {

	offerId := c.Params.ByName("offerId")

	offer, error := controller.usecase.OfferUsecase.GetOfferById(offerId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, offer)
}

func (controller OfferController) AddOffer(c *gin.Context) {

	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	var offer contract.OfferDTO

	if err := c.BindJSON(&offer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, err := controller.usecase.SearchUsecase.GetSearchById(searchId)
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	offer.UserID = userDTO.Id
	offer.SearchID = searchDTO.Id

	offerDTO, err := controller.usecase.OfferUsecase.AddOffer(offer)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, offerDTO)
}

func (controller OfferController) ModifyOffer(c *gin.Context) {
	var offer contract.OfferDTO

	if err := c.BindJSON(&offer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	offerDTO, error := controller.usecase.OfferUsecase.ModifyOffer(offer)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, offerDTO)
}
