package repository

import (
	company_repository "github.com/philvc/jobbi-api/repository/company"
	network_repository "github.com/philvc/jobbi-api/repository/network"
	offer_repository "github.com/philvc/jobbi-api/repository/offer"
	search_repository "github.com/philvc/jobbi-api/repository/search"
	user_repository "github.com/philvc/jobbi-api/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository       user_repository.UserRepository
	SearchRepository     search_repository.SearchRepository
	OfferRepository      offer_repository.OfferRepository
	CompanyRepository    company_repository.CompanyRepository
	NetworkRepository    network_repository.NetworkRepository
}

func Default(database *gorm.DB) Repository {
	return Repository{
		UserRepository:       user_repository.Default(database),
		SearchRepository:     search_repository.Default(database),
		OfferRepository:      offer_repository.Default(database),
		CompanyRepository:    company_repository.Default(database),
		NetworkRepository:    network_repository.Default(database),
	}
}
