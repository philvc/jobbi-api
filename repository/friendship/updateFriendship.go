package friendship_repository

import (
	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)

func (repository FriendshipRepository) UpdateFriendship(friendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	friendship := model.ToFriendship(friendshipDTO)

	repository.database.Model(&friendship).Where("id = ?", friendship.ID).Updates(map[string]interface{}{"state": friendship.State,
		"user_id": friendship.UserID, "search_id": friendship.SearchID})

	friendshipDTO = model.ToFriendshipDTO(friendship)

	return &friendshipDTO, nil
}
