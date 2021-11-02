package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type Offer struct {
	Base
	Description string
	Title       string
	Link        string
	SearchID    string
	UserID      string
}

func ToOfferDTO(Offer Offer) contract.OfferDTO {
	return contract.OfferDTO{
		Id:          Offer.Base.ID,
		Description: Offer.Description,
		Title:       Offer.Title,
		Link:        Offer.Link,
		SearchID:    Offer.SearchID,
		UserID:      Offer.UserID,
	}
}

func ToOffer(offerDTO contract.OfferDTO) Offer {
	return Offer{
		Base: Base{
			ID: offerDTO.Id,
		},
		Link:        offerDTO.Link,
		Description: offerDTO.Description,
		Title:       offerDTO.Title,
		SearchID:    offerDTO.SearchID,
		UserID:      offerDTO.UserID,
	}
}

func ToOfferDTOs(Offers []Offer) []contract.OfferDTO {
	OfferDTOs := make([]contract.OfferDTO, len(Offers))

	for i, item := range Offers {
		OfferDTOs[i] = ToOfferDTO(item)
	}

	return OfferDTOs
}

func ToOfferes(OffersDTO []contract.OfferDTO) []Offer {
	Offers := make([]Offer, len(OffersDTO))

	for i, item := range OffersDTO {
		Offers[i] = ToOffer(item)
	}

	return Offers
}
