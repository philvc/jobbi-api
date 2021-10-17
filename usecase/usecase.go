package usecase

import (
	"github.com/philvc/jobbi-api/repository"
	client_usecase "github.com/philvc/jobbi-api/usecase/client"
	invoice_usecase "github.com/philvc/jobbi-api/usecase/invoice"
	invoice_export_usecase "github.com/philvc/jobbi-api/usecase/invoice/export"
	organisation_usecase "github.com/philvc/jobbi-api/usecase/organisation"
	user_usecase "github.com/philvc/jobbi-api/usecase/user"
	search_usecase "github.com/philvc/jobbi-api/usecase/search"
	offer_usecase "github.com/philvc/jobbi-api/usecase/offer"
)


type Usecase struct {
	ClientUsecase       client_usecase.ClientUsecase
	InvoiceUsecase      invoice_usecase.InvoiceUsecase
	ExportUsecase       invoice_export_usecase.ExportUsecase
	OrganisationUsecase organisation_usecase.OrganisationUsecase
	UserUsecase         user_usecase.UserUsecase
	SearchUsecase 		search_usecase.SearchUseCase
	OfferUsecase        offer_usecase.OfferUseCase
}

func Default(repository repository.Repository) Usecase {
	return Usecase{
		ClientUsecase:       client_usecase.Default(repository),
		InvoiceUsecase:      invoice_usecase.Default(repository),
		ExportUsecase:       invoice_export_usecase.Default(repository),
		OrganisationUsecase: organisation_usecase.Default(repository),
		UserUsecase:         user_usecase.Default(repository),
		SearchUsecase:        search_usecase.Default(repository),
		OfferUsecase:         offer_usecase.Default(repository),
	}
}
