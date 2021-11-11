package friendship_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
)
	

func(controller FriendshipController) UpdateFriendship(c *gin.Context){
	
	// Params
	body := contract.FriendshipDTO{}

	err := c.BindJSON(&body)

	if err != nil {
		
		c.JSON(http.StatusBadRequest, err) 	 
		return 
	}

	friendship, err := controller.usecase.FriendshipUsecase.UpdateFriendship(&body)

	if err != nil {
		
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	c.JSON(http.StatusOK, friendship)

}