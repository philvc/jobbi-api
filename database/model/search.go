package model

import (
	"github.com/nightborn-be/invoice-backend/contract"
	"gorm.io/gorm"
)

type Search struct {
	gorm.Model
	Description     string
	Title 			string
	Owner	string
	Answers []Answer`gorm:"one2many:search_answers"`
}

func ToSearchDTO(search Search) contract.SearchDTO {
	return contract.SearchDTO{
		Id:         search.ID,
		Description:  search.Description,
		Title: search.Title,
		Owner: search.Owner,
	}
}

func ToSearch(searchDTO contract.SearchDTO) Search {
	return Search{
		Model: gorm.Model{
			ID: searchDTO.Id,
		},
		Description:  searchDTO.Description,
		Title: searchDTO.Title,
		Owner: searchDTO.Owner,
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
