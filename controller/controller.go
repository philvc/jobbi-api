package controller

import (
	client_controller "github.com/philvc/jobbi-api/controller/client"
	client_invoice_controller "github.com/philvc/jobbi-api/controller/client/invoice"
	company_controller "github.com/philvc/jobbi-api/controller/company"
	invoice_controller "github.com/philvc/jobbi-api/controller/invoice"
	invoice_export_controller "github.com/philvc/jobbi-api/controller/invoice/export"
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
	ClientController           client_controller.ClientController
	ClientInvoiceController    client_invoice_controller.ClientInvoiceController
	InvoiceController          invoice_controller.InvoiceController
	ExportController           invoice_export_controller.ExportController
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
		ClientController:           client_controller.Default(usecase),
		InvoiceController:          invoice_controller.Default(usecase),
		ClientInvoiceController:    client_invoice_controller.Default(usecase),
		ExportController:           invoice_export_controller.Default(usecase),
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
