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

func (repository FriendshipRepository) GetFriendshipsByState(userId uint, state uint ) (*[]contract.FriendshipDTO, error){
	var friendships []model.Friendship
	var user model.User

	if err := repository.database.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Where("State = ? AND UserID = ?", state, userId).Find(&friendships).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}
func (repository FriendshipRepository) GetFriendshipsBySearchId(searchId uint) (*[]contract.FriendshipDTO, error){
	var friendships []model.Friendship
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Where("State = ? AND SearchID = ?", 1, searchId).Find(&friendships).Error; err != nil {
		return nil, errors.New(err.Error())
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

	repository.database.Model(&friendship).Where("id = ?", friendship.ID).Updates(map[string]interface{}{"State": friendship.State})

	friendshipDTO = model.ToFriendshipDTO(friendship)

	return &friendshipDTO, nil
}
