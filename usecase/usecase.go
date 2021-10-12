package usecase

import (
	"github.com/nightborn-be/invoice-backend/repository"
	client_usecase "github.com/nightborn-be/invoice-backend/usecase/client"
	invoice_usecase "github.com/nightborn-be/invoice-backend/usecase/invoice"
	invoice_export_usecase "github.com/nightborn-be/invoice-backend/usecase/invoice/export"
	organisation_usecase "github.com/nightborn-be/invoice-backend/usecase/organisation"
	user_usecase "github.com/nightborn-be/invoice-backend/usecase/user"
)

type Usecase struct {
	ClientUsecase       client_usecase.ClientUsecase
	InvoiceUsecase      invoice_usecase.InvoiceUsecase
	ExportUsecase       invoice_export_usecase.ExportUsecase
	OrganisationUsecase organisation_usecase.OrganisationUsecase
	UserUsecase         user_usecase.UserUsecase
}

func Default(repository repository.Repository) Usecase {
	return Usecase{
		ClientUsecase:       client_usecase.Default(repository),
		InvoiceUsecase:      invoice_usecase.Default(repository),
		ExportUsecase:       invoice_export_usecase.Default(repository),
		OrganisationUsecase: organisation_usecase.Default(repository),
		UserUsecase:         user_usecase.Default(repository),
	}
}
