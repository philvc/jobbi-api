package search_repository

import (
	"errors"

	"github.com/google/uuid"
	contract "github.com/philvc/jobbi-api/contract"
	constant "github.com/philvc/jobbi-api/constants"
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

func (repository SearchRepository) GetMySearch(userId string) (*contract.MySearchDTO, error) {

	var result contract.MySearchDTO

	// Get search by user id
	if err := repository.database.
	Model(&model.Search{}).
	Select("title, id, sector, tags").
	Where("user_id = ? ", userId).
	Scan(&result).
	Error; err != nil {
		print(err)
		return nil, err
	}

	// If user has a search
	if result.Id != "" {

		// Get participants
		if err := repository.database.Model(&model.Friendship{}).
			Where("search_id = ? ", result.Id).
			Joins("JOIN users ON users.id = friendships.user_id").
			Select("users.id, users.first_name, users.last_name, users.avatar_url").
			Find(&result.Participants).Error; err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (repository SearchRepository) GetSharedSearches(userId string)(*[]contract.SharedSearchDTO , error){
	
	var results []contract.SharedSearchDTO

	if err := repository.database.
	Model(&model.Friendship{}).
	Where("friendships.user_id = ?", userId).
	
	// Get friendships where type is INVITED
	Where("friendships.type = ?", constant.FRIENDSHIP_TYPE_INVITED).
	
	// GET Searches details
	Joins("JOIN searches ON searches.id = friendships.search_id").
	
	// Get search owner details
	Joins("JOIN users ON users.id = searches.user_id").
	Select("searches.id, searches.title, searches.sector, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
	Find(&results).
	Error; err != nil {
		return nil, err
	}
	
	return &results, nil
}

func (repository SearchRepository) GetFollowedSearches(userId string)(*[]contract.FollowedSearchDTO, error){

	var results []contract.FollowedSearchDTO

	if err := repository.database.
	Model(&model.Friendship{}).
	Where("friendships.user_id = ?", userId).
	
	// Get friendships where type is INVITED
	Where("friendships.type = ?", constant.FRIENDSHIP_TYPE_FOLLOWED).
	
	// GET Searches details
	Joins("JOIN searches ON searches.id = friendships.search_id").
	
	// Get search owner details
	Joins("JOIN users ON users.id = searches.user_id").
	Select("searches.id, searches.sector, searches.title, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
	Find(&results).
	Error; err != nil {
		return nil, err
	}
	
	return &results, nil
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

	id := uuid.New()

	search.ID = id.String()

	if err := repository.database.Create(&search).Error; err != nil {
		return nil, errors.New("failed to create Search")
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) ModifySearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	repository.database.Model(&search).Where("id = ?", search.ID).Updates(map[string]interface{}{"title": search.Title,
		"description": search.Description})

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}
