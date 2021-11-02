package offer_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type OfferRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) OfferRepository {
	return OfferRepository{
		database: db,
	}
}

func (repository OfferRepository) GetOffersBySearchId(searchId string) (*[]contract.OfferDTO, error) {
	var Offers []model.Offer
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&search).Association("Offers").Find(&Offers); err != nil {
		return nil, err
	}

	OfferDTOs := model.ToOfferDTOs(Offers)

	return &OfferDTOs, nil
}

func (repository OfferRepository) GetOfferById(OfferId string) (*contract.OfferDTO, error) {
	var Offer model.Offer

	if err := repository.database.Where("id = ?", OfferId).First(&Offer).Error; err != nil {
		return nil, errors.New("no Offer")
	}

	OfferDTO := model.ToOfferDTO(Offer)

	return &OfferDTO, nil
}

func (repository OfferRepository) AddOffer(OfferDTO contract.OfferDTO) (*contract.OfferDTO, error) {

	Offer := model.ToOffer(OfferDTO)

	if err := repository.database.Create(&Offer).Error; err != nil {
		return nil, errors.New("failed to create Offer")
	}

	OfferDTO = model.ToOfferDTO(Offer)

	return &OfferDTO, nil
}

func (repository OfferRepository) ModifyOffer(OfferDTO contract.OfferDTO) (*contract.OfferDTO, error) {

	Offer := model.ToOffer(OfferDTO)

	repository.database.Model(&Offer).Where("id = ?", Offer.ID).Updates(map[string]interface{}{
		"link": Offer.Link, "description": Offer.Description, "title": Offer.Title})

	OfferDTO = model.ToOfferDTO(Offer)

	return &OfferDTO, nil
}

func (repository OfferRepository) DeleteOffer(offerId string) (bool, error) {

	var Offer model.Offer

	if err := repository.database.Where("id = ?", offerId).First(&Offer).Error; err != nil {
		return false, errors.New("no Offer")
	}

	if err := repository.database.Delete(&Offer, offerId).Error; err != nil {
		return false, errors.New("delete offer failed")
	}

	return true, nil
}
