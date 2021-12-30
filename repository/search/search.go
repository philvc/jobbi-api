package search_repository

import (
	"errors"

	"github.com/google/uuid"
	constant "github.com/philvc/jobbi-api/constants"
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

func (repository SearchRepository) GetMySearch(userId string) (*contract.MySearchDTO, error) {

	var result contract.MySearchDTO

	// Get search by user id
	if err := repository.database.
		Model(&model.Search{}).
		Select("title, id, sector, tags, type").
		Where("user_id = ? ", userId).
		Scan(&result).
		Error; err != nil {
		print(err)
		return nil, errors.New(constant.ErrorGetMySearch)
	}

	// If user has a search
	if result.Id != "" {

		// Get participants
		if err := repository.database.Model(&model.Friendship{}).
			Where("search_id = ? ", result.Id).
			Joins("JOIN users ON users.id = friendships.user_id").
			Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url").
			Find(&result.Participants).Error; err != nil {
			return nil, errors.New(constant.ErrorGetMySearchParticipants)
		}
	}

	return &result, nil
}

func (repository SearchRepository) GetSharedSearches(userId string) (*[]contract.SharedSearchDTO, error) {

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
		Select("searches.id, searches.title, searches.type, searches.sector, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetSharedSearches)
	}

	return &results, nil
}

func (repository SearchRepository) GetFollowedSearches(userId string) (*[]contract.FollowedSearchDTO, error) {

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
		Select("searches.id, searches.sector, searches.type, searches.title, searches.tags, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetFollowedSearches)
	}

	return &results, nil
}

func (repository SearchRepository) GetSearchById(searchId string) (*contract.SearchDTOById, error) {
	var result contract.SearchDTOById

	if err := repository.database.
		Model(&model.Search{}).
		Where("searches.id = ?", searchId).
		Select("title, searches.id, sector, tags, description, searches.type, users.email, users.id as user_id, users.first_name, users.last_name, users.avatar_url").
		Joins("JOIN users ON users.id = searches.user_id").
		Scan(&result).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetSearchById)
	}

	// If user has a search
	if result.Id != "" {

		// Get participants
		if err := repository.database.Model(&model.Friendship{}).
			Where("search_id = ? ", result.Id).
			Joins("JOIN users ON users.id = friendships.user_id").
			Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url").
			Find(&result.Participants).Error; err != nil {
			return nil, errors.New(constant.ErrorGetMySearchParticipants)
		}
	}

	return &result, nil
}

func (repository SearchRepository) GetPostsBySearchId(searchId string) (*[]contract.PostDTOBySearchId, error) {

	var results []contract.PostDTOBySearchId

	if err := repository.database.
		Model(&model.Post{}).
		Where("search_id = ?", searchId).
		Joins("JOIN users ON users.id = posts.user_id").
		Select("users.id, posts.id, users.email, users.first_name as user_first_name, users.last_name as user_last_name, title").
		Find(&results).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPostsBySearchId)
	}

	return &results, nil

}

func (repository SearchRepository) GetParticipantsBySearchId(searchId string) (*[]contract.ParticipantDTOForSearchById, error) {

	var results []contract.ParticipantDTOForSearchById

	var posts []model.Post

	if err := repository.database.
		Model(&model.Friendship{}).
		Where("friendships.search_id = ?", searchId).
		Joins("JOIN users ON users.id = friendships.user_id").
		Select("users.id, users.first_name, users.last_name, users.email, users.avatar_url, friendships.type").
		Find(&results).
		Joins("JOIN posts ON posts.user_id = users.id").
		Select("posts.id, posts.user_id").
		Find(&posts).
		Error; err != nil {
		return nil, errors.New(constant.ErrorGetPostsBySearchId)
	}

	// Count total of post for each participant
	if len(results) != 0 {
		for index, item := range results {
			var count int64 = 0

			for _, post := range posts {
				if post.UserID == item.Id {
					count = +1
				}
			}

			results[index].NumberOfPosts = count
		}
	}

	return &results, nil

}

func (repository SearchRepository) AddSearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	// Add new search uuid
	id := uuid.New()

	search.ID = id.String()

	if err := repository.database.Create(&search).Error; err != nil {
		return nil, errors.New(constant.ErrorAddSearch)
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) ModifySearch(SearchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	search := model.ToSearch(SearchDTO)

	if err := repository.database.Model(&search).Where("id = ?", search.ID).Updates(map[string]interface{}{"title": search.Title,
		"description": search.Description}).Error; err != nil {
		return nil, errors.New(constant.ErrorModifySearch)
	}

	searchDTO := model.ToSearchDTO(search)

	return &searchDTO, nil
}

func (repository SearchRepository) IsSearchOwner(userId string, searchId string) bool {

	if err := repository.database.Model(&model.Search{}).Where("id = ?", searchId).Where("user_id = ?", userId).First(&model.Search{}).Error; err != nil {
		return false
	}

	return true
}
func (repository SearchRepository) IsPublic(searchId string) bool {

	var search model.Search

	// Get search by id
	if err := repository.database.Model(&model.Search{}).Where("id = ?", searchId).Where("type = ?", constant.SearchTypePublic).First(&search).Error; err != nil {
		return false
	}

	return true
}

func (repository SearchRepository) IsFriend(userId string, searchId string) bool {

	var friendship model.Friendship

	// Get search by id
	if err := repository.database.Model(&model.Friendship{}).Where("search_id = ?", searchId).Where("user_id = ? ", userId).First(&friendship).Error; err != nil {
		return false

	}

	return true
}
