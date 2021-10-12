package repository

import (
	client_repository "github.com/nightborn-be/invoice-backend/repository/client"
	invoice_repository "github.com/nightborn-be/invoice-backend/repository/invoice"
	organisation_repository "github.com/nightborn-be/invoice-backend/repository/organisation"
	user_repository "github.com/nightborn-be/invoice-backend/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository         user_repository.UserRepository
	ClientRepository       client_repository.ClientRepository
	InvoiceRepository      invoice_repository.InvoiceRepository
	OrganisationRepository organisation_repository.OrganisationRepository
}

func Default(database *gorm.DB) Repository {
	return Repository{
		UserRepository:         user_repository.Default(database),
		ClientRepository:       client_repository.Default(database),
		InvoiceRepository:      invoice_repository.Default(database),
		OrganisationRepository: organisation_repository.Default(database),
	}
}
