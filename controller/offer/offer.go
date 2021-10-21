package offer_controller

import (
	"net/http"
	"strconv"

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

// swagger:operation GET /searches/{searchId}/offers offers GetOffersBySearchId
// type id struct
// Get offers by searchId.
// Return offer
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
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
//             $ref: "#/definitions/OfferDTO"
//       400:
//         description: Bad Request

func (controller OfferController) GetOffersBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	offers, error := controller.usecase.OfferUsecase.GetOffersBySearchId(searchId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, offers)
}

// swagger:operation GET /searches/{searchId}/offers/{offerId} offers GetOfferById
// type id struct
// Get offer by id.
// Return offer
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: offerId
//         in: path
//         type: string
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/OfferDTO"
//       400:
//         description: Bad Request
func (controller OfferController) GetOfferById(c *gin.Context) {

	offerId := c.Params.ByName("offerId")

	offer, error := controller.usecase.OfferUsecase.GetOfferById(offerId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, offer)
}

// swagger:operation POST /searches/{searchId}/offers offers AddOffer
// type id struct
// Create offer.
// Return offer
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: offer
//         in: body
//         schema:
//            $ref: "#/definitions/OfferDTO"
//         description: offer
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/OfferDTO"
//       400:
//         description: Bad Request
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

// swagger:operation PUT /searches/{searchId}/offers/{offerId} offers ModifyOffer
// type id struct
// Modify offer.
// Return offer
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: offerId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: offer
//         in: body
//         schema:
//            $ref: "#/definitions/OfferDTO"
//         description: offer
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/OfferDTO"
//       400:
//         description: Bad Request
func (controller OfferController) ModifyOffer(c *gin.Context) {

	offerId := c.Params.ByName(("offerId"))

	var offer contract.OfferDTO

	if err := c.BindJSON(&offer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	parsedId, _ := strconv.ParseUint(offerId, 10, 32)
	offer.Id = uint(parsedId)

	offerDTO, error := controller.usecase.OfferUsecase.ModifyOffer(offer)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, offerDTO)
}

// swagger:operation DELETE /searches/{searchId}/offers/{offerId} offers DeleteOffer
// type id struct
// Delete offer.
// Return true or error
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: offerId
//         in: path
//         type: string
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//           type: boolean
//       400:
//         description: Bad Request
func (controller OfferController) DeleteOffer(c *gin.Context) {
	offerId := c.Params.ByName(("offerId"))

	var offer contract.OfferDTO

	parsedId, _ := strconv.ParseUint(offerId, 10, 32)

	offer.Id = uint(parsedId)

	result, error := controller.usecase.OfferUsecase.DeleteOffer(offer.Id)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}


