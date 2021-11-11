package friendship_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
)

func (u FriendshipUsecase) CreateFriendship(friendshipDto *contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	// check missing fields
	if friendshipDto.SearchId == "" || friendshipDto.UserId == "" {
		return nil, errors.New("missing-friendship-field")
	}

	// Post friendship
	friendship, err := u.repository.FriendshipRepository.CreateFriendship(*friendshipDto)

	if err != nil {
		return nil, err
	}

	return friendship, nil

}
