package usecase

import (
	"github.com/philvc/jobbi-api/repository"
	company_usecase "github.com/philvc/jobbi-api/usecase/company"
	network_usecase "github.com/philvc/jobbi-api/usecase/network"
	offer_usecase "github.com/philvc/jobbi-api/usecase/offer"
	search_usecase "github.com/philvc/jobbi-api/usecase/search"
	user_usecase  "github.com/philvc/jobbi-api/usecase/user"
)

type Usecase struct {
	UserUsecase       user_usecase.UserUsecase
	SearchUsecase     search_usecase.SearchUseCase
	OfferUsecase      offer_usecase.OfferUseCase
	CompanyUsecase    company_usecase.CompanyUseCase
	NetworkUsecase    network_usecase.NetworkUseCase
}

func Default(repository repository.Repository) Usecase {
	return Usecase{
		UserUsecase:       user_usecase.Default(repository),
		SearchUsecase:     search_usecase.Default(repository),
		OfferUsecase:      offer_usecase.Default(repository),
		CompanyUsecase:    company_usecase.Default(repository),
		NetworkUsecase:    network_usecase.Default(repository),
	}
}
