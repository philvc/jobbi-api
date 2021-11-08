package friendship_repository

import (
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)

func (repo FriendshipRepositoryImplementation) GetFriendshipsByUserIdAndState(userId string, state uint) (*[]contract.FriendshipDTO, error) {
	var friendships []model.Friendship

	if err := repo.database.Where("user_id = ? AND state = ?", userId, state).Find(&friendships).Error; err != nil {
		return nil, err
	}

	friendshipDTOs := model.ToFriendshipDTOs(friendships)

	return &friendshipDTOs, nil
}
