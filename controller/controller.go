package controller

import (
	company_controller "github.com/philvc/jobbi-api/controller/company"
	network_controller "github.com/philvc/jobbi-api/controller/network"
	offer_controller "github.com/philvc/jobbi-api/controller/offer"
	search_controller "github.com/philvc/jobbi-api/controller/search"
	user_controller "github.com/philvc/jobbi-api/controller/user"
	friendship_controller "github.com/philvc/jobbi-api/controller/friendship"

	"github.com/philvc/jobbi-api/usecase"
)

type Controller struct {
	UserController             user_controller.UserController
	SearchController           search_controller.SearchController
	OfferController            offer_controller.OfferController
	CompanyController          company_controller.CompanyController
	NetworkController          network_controller.NetworkController
	FriendshipController       friendship_controller.FriendshipController
}

func Default(usecase usecase.Usecase) Controller {
	return Controller{
		UserController:             user_controller.Default(usecase),
		SearchController:           search_controller.Default(usecase),
		OfferController:            offer_controller.Default(usecase),
		CompanyController:          company_controller.Default(usecase),
		NetworkController:          network_controller.Default(usecase),
		FriendshipController:       friendship_controller.Default(usecase),
	}
}
