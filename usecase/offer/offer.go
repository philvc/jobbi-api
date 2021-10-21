package offer_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type OfferUseCase struct {
	repository repository.Repository
}

// Returns an instance of a offer use-case
func Default(repository repository.Repository) OfferUseCase {
	return OfferUseCase{
		repository: repository,
	}
}

func (usecase OfferUseCase) GetOffersBySearchId(searchId string) (*[]contract.OfferDTO, error) {
	offers, err := usecase.repository.OfferRepository.GetOffersBySearchId(searchId)
	return offers, err
}

func (usecase OfferUseCase) GetOfferById(offerId string) (*contract.OfferDTO, error) {
	offer, err := usecase.repository.OfferRepository.GetOfferById(offerId)
	return offer, err
}

func (usecase OfferUseCase) AddOffer(OfferDTO contract.OfferDTO) (*contract.OfferDTO, error) {

	if OfferDTO.Title == "" {
		return nil, errors.New("missing title")
	}

	if OfferDTO.Description == "" {
		return nil, errors.New("missing description")
	}
	if OfferDTO.Link == "" {
		return nil, errors.New("missing link")
	}
	offer, err := usecase.repository.OfferRepository.AddOffer(OfferDTO)
	return offer, err
}

func (usecase OfferUseCase) ModifyOffer(OfferDTO contract.OfferDTO) (*contract.OfferDTO, error) {
	user, err := usecase.repository.OfferRepository.ModifyOffer(OfferDTO)
	return user, err
}

func (usecase OfferUseCase) DeleteOffer(offerId uint) (bool, error) {
	result, err := usecase.repository.OfferRepository.DeleteOffer(offerId)
	return result, err
}
