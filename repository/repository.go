package repository

import (
	client_repository "github.com/philvc/jobbi-api/repository/client"
	company_repository "github.com/philvc/jobbi-api/repository/company"
	friendship_repository "github.com/philvc/jobbi-api/repository/friendship"
	invoice_repository "github.com/philvc/jobbi-api/repository/invoice"
	network_repository "github.com/philvc/jobbi-api/repository/network"
	offer_repository "github.com/philvc/jobbi-api/repository/offer"
	organisation_repository "github.com/philvc/jobbi-api/repository/organisation"
	search_repository "github.com/philvc/jobbi-api/repository/search"
	user_repository "github.com/philvc/jobbi-api/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository         user_repository.UserRepository
	ClientRepository       client_repository.ClientRepository
	InvoiceRepository      invoice_repository.InvoiceRepository
	OrganisationRepository organisation_repository.OrganisationRepository
	SearchRepository       search_repository.SearchRepository
	OfferRepository        offer_repository.OfferRepository
	CompanyRepository      company_repository.CompanyRepository
	NetworkRepository      network_repository.NetworkRepository
	FriendshipRepository   friendship_repository.FriendshipRepository
}

func Default(database *gorm.DB) Repository {
	return Repository{
		UserRepository:         user_repository.Default(database),
		ClientRepository:       client_repository.Default(database),
		InvoiceRepository:      invoice_repository.Default(database),
		OrganisationRepository: organisation_repository.Default(database),
		SearchRepository:       search_repository.Default(database),
		OfferRepository:        offer_repository.Default(database),
		CompanyRepository:      company_repository.Default(database),
		NetworkRepository:      network_repository.Default(database),
		FriendshipRepository:   friendship_repository.Default(database),
	}
}
