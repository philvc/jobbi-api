package search_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller SearchController) GetSearchFriendships(c *gin.Context) {


	// Params
	searchId := c.Params.ByName("searchId")
	status := c.Query("status")

	parseStatus, _ := strconv.ParseInt(status, 10,32)
	intStatus := int(parseStatus)

	results, err := controller.usecase.SearchUsecase.GetSearchFriendships(searchId, intStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}
