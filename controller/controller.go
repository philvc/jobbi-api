package controller

import (
	company_controller "github.com/philvc/jobbi-api/controller/company"
	network_controller "github.com/philvc/jobbi-api/controller/network"
	offer_controller "github.com/philvc/jobbi-api/controller/offer"
	organisation_controller "github.com/philvc/jobbi-api/controller/organisation"
	search_controller "github.com/philvc/jobbi-api/controller/search"
	user_controller "github.com/philvc/jobbi-api/controller/user"
	user_organisation_controller "github.com/philvc/jobbi-api/controller/user/organisation"
	friendship_controller "github.com/philvc/jobbi-api/controller/friendship"

	"github.com/philvc/jobbi-api/usecase"
)

type Controller struct {
	OrganisationController     organisation_controller.OrganisationController
	UserController             user_controller.UserController
	UserOrganisationController user_organisation_controller.UserOrganisationController
	SearchController           search_controller.SearchController
	OfferController            offer_controller.OfferController
	CompanyController          company_controller.CompanyController
	NetworkController          network_controller.NetworkController
	FriendshipController       friendship_controller.FriendshipController
}

func Default(usecase usecase.Usecase) Controller {
	return Controller{
		OrganisationController:     organisation_controller.Default(usecase),
		UserController:             user_controller.Default(usecase),
		UserOrganisationController: user_organisation_controller.Default(usecase),
		SearchController:           search_controller.Default(usecase),
		OfferController:            offer_controller.Default(usecase),
		CompanyController:          company_controller.Default(usecase),
		NetworkController:          network_controller.Default(usecase),
		FriendshipController:       friendship_controller.Default(usecase),
	}
}
