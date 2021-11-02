package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type Search struct {
	Base
	Description string
	Title       string
	UserID      string
	Friendships []Friendship `gorm:"foreignKey:SearchID"`
	Offers      []Offer      `gorm:"foreignKey:SearchID"`
	Companies   []Company    `gorm:"foreignKey:SearchID"`
	Networks    []Network    `gorm:"foreignKey:SearchID"`
}

func ToSearchDTO(search Search) contract.SearchDTO {
	return contract.SearchDTO{
		Id:          search.Base.ID,
		Description: search.Description,
		Title:       search.Title,
		UserID:      search.UserID,
	}
}

func ToSearch(searchDTO contract.SearchDTO) Search {
	return Search{
		Base: Base{
			ID: searchDTO.Id,
		},
		Description: searchDTO.Description,
		Title:       searchDTO.Title,
		UserID:      searchDTO.UserID,
	}
}

func ToSearchDTOs(searches []Search) []contract.SearchDTO {
	SearchDTOs := make([]contract.SearchDTO, len(searches))

	for i, item := range searches {
		SearchDTOs[i] = ToSearchDTO(item)
	}

	return SearchDTOs
}

func ToSearches(searchesDTO []contract.SearchDTO) []Search {
	Searches := make([]Search, len(searchesDTO))

	for i, item := range searchesDTO {
		Searches[i] = ToSearch(item)
	}

	return Searches
}
