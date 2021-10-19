package friendship_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/usecase"
)

type FriendshipController struct {
	usecase usecase.Usecase
}

func Default(usecase usecase.Usecase) FriendshipController {
	return FriendshipController{
		usecase: usecase,
	}
}

func (controller FriendshipController) GetFriendshipsBySearchId(c *gin.Context) {
	searchId := c.Params.ByName("searchId")

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySearchId(searchId)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Friendships)
}
func (controller FriendshipController) GetFriendshipsBySub(c *gin.Context) {
	sub := c.GetString("sub")

	Friendships, error := controller.usecase.FriendshipUsecase.GetFriendshipsBySub(sub)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error)
		return
	}

	c.IndentedJSON(http.StatusOK, Friendships)
}

func (controller FriendshipController) AddFriendship(c *gin.Context) {

	searchId := c.Params.ByName("searchId")

	var Friendship contract.FriendshipDTO

	if err := c.BindJSON(&Friendship); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	searchDTO, err := controller.usecase.SearchUsecase.GetSearchById(searchId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	Friendship.SearchId = searchDTO.Id

	FriendshipDTO, err := controller.usecase.FriendshipUsecase.AddFriendship(Friendship)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, FriendshipDTO)
}

func (controller FriendshipController) ModifyFriendship(c *gin.Context) {
	var Friendship contract.FriendshipDTO

	if err := c.BindJSON(&Friendship); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	FriendshipDTO, error := controller.usecase.FriendshipUsecase.ModifyFriendship(Friendship)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, FriendshipDTO)
}
