package friendship_usecase

import "github.com/philvc/jobbi-api/contract"


func (u FriendshipUsecase)UpdateFriendship(friendshipDto *contract.FriendshipDTO)( *contract.FriendshipDTO, error){

	friendship, err := u.repository.FriendshipRepository.UpdateFriendship(*friendshipDto)

	if err != nil {
		return nil, err
	}

	return friendship, nil
}
