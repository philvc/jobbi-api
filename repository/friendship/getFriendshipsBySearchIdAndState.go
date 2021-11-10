package friendship_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)

func (repository FriendshipRepository) GetFriendshipsBySearchIdAndState(searchId string, status uint) (*[]contract.FriendshipDTO, error) {
	var friendships []model.Friendship

	if err := repository.database.Where("search_id = ? AND state = ?", searchId, status).Find(&friendships).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}
