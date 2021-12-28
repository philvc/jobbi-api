package search_repository

import (
	"errors"

	"github.com/google/uuid"
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
	Select("title, id, tags").
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
	
	// friendships where user id & joins searches and joins user and scan
	var results []contract.SharedSearchDTO

	if err := repository.database.
	Model(&model.Friendship{}).
	Where("friendships.user_id = ?", userId).
	Joins("JOIN searches ON searches.id = friendships.search_id").
	Joins("JOIN users ON users.id = searches.user_id").
	Select("searches.id, searches.tags, searches.title, searches.description, searches.user_id, users.first_name, users.last_name, users.avatar_url").
	Find(&results).
	Error; err != nil {
		return nil, err
	}
	
	return &results, nil
}

func (repository SearchRepository) GetFriendsSearches(userId string) (*[]contract.FriendSearchDTO, error) {

	// Get user Friendships then fetch search owner
	// TODO CONTINUE MODEL & the return statementTHEN MAPPER USECASE
	var friendsrhips []model.Friendship
	type searchTitle struct {
		Title     string
		FirstName string
	}

	var results []searchTitle
	var searches []model.Search

	if err := repository.database.Model(&model.Search{}).
		Select("title, users.first_name").
		Joins("JOIN friendships ON friendships.search_id = searches.id").
		Joins("JOIN users ON users.id = searches.user_id").
		Where("friendships.user_id = ?", userId).
		Take(&searches).
		Scan(&results).
		Error; err != nil {
		print(err)
	}

	if err := repository.database.Find(&friendsrhips).Error; err != nil {
		return nil, err
	}

	return nil, nil
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
