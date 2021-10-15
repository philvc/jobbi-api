package search_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type SearchRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) SearchRepository {
	return SearchRepository{
		database: db,
	}
}

func (repository SearchRepository) GetSearchesByUserId(userId uint) (*[]contract.SearchDTO, error) {
	var searches []model.Search
	var user model.User

	if err := repository.database.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&user).Association("Searches").Find(&searches); err != nil {
		return nil, err
	}

	searchDTOs := model.ToSearchDTOs(searches)

	return &searchDTOs, nil
}

func (repository SearchRepository) GetSearchById(searchId string) (*contract.SearchDTO, error) {
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New("no Search")
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) AddSearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	if err := repository.database.Create(&search).Error; err != nil {
		return nil, errors.New("failed to create Search")
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) ModifySearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	repository.database.Model(&search).Where("id = ?", search.ID).Updates(map[string]interface{}{"title": search.Title,
		"description": search.Description, "owner": search.Owner})

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}
