package user_usecase

import (
	contract "github.com/philvc/jobbi-api/contract/search"
)




func (u UserUsecase) GetFriendshipsByUserIdAndState(searchId string, state uint) (*contract.SearchWithOwnerAndFriends, error) {

	dto := &contract.SearchWithOwnerAndFriends{}

	u.repository.FriendshipRepository.GetFriendshipsByUserIdAndState(searchId, state)
	return dto, nil
}
