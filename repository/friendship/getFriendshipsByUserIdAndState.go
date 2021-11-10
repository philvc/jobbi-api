package friendship_repository

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)

func (repo FriendshipRepository) GetFriendshipsByUserIdAndState(userId string, state int) (*[]contract.FriendshipDTO, error) {
	
	friendships := []model.Friendship{}

	if err := repo.database.Where("user_id = ?", userId).Find(&friendships).Error; err != nil {
		return nil, errors.New("error-fetching-friendships")
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}
