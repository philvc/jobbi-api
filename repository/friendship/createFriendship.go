package friendship_repository

import (
	"errors"

	"github.com/google/uuid"
	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)


func (repository FriendshipRepository) CreateFriendship(friendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	friendship := model.ToFriendship(friendshipDTO)

	id := uuid.New()

	friendship.ID = id.String()

	if err := repository.database.Create(&friendship).Error; err != nil {
		return nil, errors.New("failed to create Friendship")
	}

	friendshipDTO = model.ToFriendshipDTO(friendship)

	return &friendshipDTO, nil
}
