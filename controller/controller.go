package controller

import (
	client_controller "github.com/philvc/jobbi-api/controller/client"
	client_invoice_controller "github.com/philvc/jobbi-api/controller/client/invoice"
	invoice_controller "github.com/philvc/jobbi-api/controller/invoice"
	invoice_export_controller "github.com/philvc/jobbi-api/controller/invoice/export"
	organisation_controller "github.com/philvc/jobbi-api/controller/organisation"
	user_controller "github.com/philvc/jobbi-api/controller/user"
	user_organisation_controller "github.com/philvc/jobbi-api/controller/user/organisation"
	search_controller "github.com/philvc/jobbi-api/controller/search"
	offer_controller "github.com/philvc/jobbi-api/controller/offer"

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
		OfferController: 			offer_controller.Default(usecase),
	}
}
