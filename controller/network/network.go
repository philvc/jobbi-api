package network_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type NetworkController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) NetworkController {
	return NetworkController{
		usecase: usecase,
	}
}

// swagger:operation GET /searches/{searchId}/networks networks GetNetworksBySearchId
// type id struct
// Get networks by searchId.
// Return network
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
//             $ref: "#/definitions/NetworkDTO"
//       400:
//         description: Bad Request
func (controller NetworkController) GetNetworksBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	Networks, error := controller.usecase.NetworkUsecase.GetNetworksBySearchId(searchId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Networks)
}

// swagger:operation GET /searches/{searchId}/networks/{networkId} networks GetNetworkById
// type id struct
// Get  network by id.
// Return  network
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name:  networkId
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
//            $ref: "#/definitions/NetworkDTO"
//       400:
//         description: Bad Request
func (controller NetworkController) GetNetworkById(c *gin.Context) {

	NetworkId := c.Params.ByName("networkId")

	Network, error := controller.usecase.NetworkUsecase.GetNetworkById(NetworkId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Network)
}


// swagger:operation POST /searches/{searchId}/networks networks AddNetwork
// type id struct
// Post  network.
// Return  network
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name:  network
//         in: body
//         schema:
//            $ref: "#/definitions/NetworkDTO"
//         description: offer
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/NetworkDTO"
//       400:
//         description: Bad Request
func (controller NetworkController) AddNetwork(c *gin.Context) {

	searchId := c.Params.ByName("searchId")
	sub := c.GetString("sub")

	var Network contract.NetworkDTO

	if err := c.BindJSON(&Network); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, err := controller.usecase.SearchUsecase.GetSearchById(searchId)
	userDTO, err := controller.usecase.UserUsecase.GetUserBySub(sub)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	Network.UserID = userDTO.Id
	Network.SearchID = searchDTO.Id

	NetworkDTO, err := controller.usecase.NetworkUsecase.AddNetwork(Network)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, NetworkDTO)
}

// swagger:operation PUT /searches/{searchId}/networks/{networkId} networks ModifyNetwork
// type id struct
// Put  network.
// Return  network
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: networkId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name:  network
//         in: body
//         schema:
//            $ref: "#/definitions/NetworkDTO"
//         description: network
//         required: true
//         description: test
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/NetworkDTO"
//       400:
//         description: Bad Request
func (controller NetworkController) ModifyNetwork(c *gin.Context) {
	var Network contract.NetworkDTO

	if err := c.BindJSON(&Network); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	NetworkDTO, error := controller.usecase.NetworkUsecase.ModifyNetwork(Network)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, NetworkDTO)
}
