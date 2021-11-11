package search_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /searches/{searchId}/friendships searches GetSearchById
// type id struct
// Get friendships by search id.
// Return users
// ---
//     Parameters:
//       - name: searchId
//         in: path
//         type: string
//         required: true
//         description: test
//       - name: status
//		   in: query
//		   type: string
//		   required: false
//     Produces:
//       - application/json
//     Responses:
//       200:
//         description: Success
//         schema:
//            $ref: "#/definitions/UserDTO"
//       400:
//         description: Bad Request
func (controller SearchController) GetSearchFriendships(c *gin.Context) {

	// Params
	searchId := c.Params.ByName("searchId")
	status := c.Query("status")

	parseStatus, _ := strconv.ParseInt(status, 10, 32)
	intStatus := int(parseStatus)

	results, err := controller.usecase.SearchUsecase.GetSearchFriendships(searchId, intStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}
