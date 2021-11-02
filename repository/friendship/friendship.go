package friendship_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type FriendshipRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) FriendshipRepository {
	return FriendshipRepository{
		database: db,
	}
}

func (repository FriendshipRepository) GetFriendshipsBySearchId(searchId string, status uint) (*[]contract.FriendshipDTO, error) {
	var friendships []model.Friendship
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Where("search_id = ? AND state = ?", search.ID, status).Find(&friendships).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}
func (repository FriendshipRepository) GetFriendshipsByUserId(userId string) (*[]contract.FriendshipDTO, error) {
	var friendships []model.Friendship
	var user model.User

	if err := repository.database.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Where("user_id = ?", userId).Association("Friendships").Find(&friendships); err != nil {
		return nil, err
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}

func (repository FriendshipRepository) AddFriendship(friendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	friendship := model.ToFriendship(friendshipDTO)

	if err := repository.database.Create(&friendship).Error; err != nil {
		return nil, errors.New("failed to create Friendship")
	}

	friendshipDTO = model.ToFriendshipDTO(friendship)

	return &friendshipDTO, nil
}

func (repository FriendshipRepository) ModifyFriendship(friendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	friendship := model.ToFriendship(friendshipDTO)

	repository.database.Model(&friendship).Where("id = ?", friendship.ID).Updates(map[string]interface{}{"state": friendship.State,
		"first_name": friendship.FirstName, "last_name": friendship.LastName, "email": friendship.Email, "user_id": friendship.UserID, "search_id": friendship.SearchID})

	friendshipDTO = model.ToFriendshipDTO(friendship)

	return &friendshipDTO, nil
}
