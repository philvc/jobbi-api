package search_usecase

import (
	"github.com/philvc/jobbi-api/contract"
)

func (u SearchUseCase) GetSearchFriendships(searchId string, state int) (*[]contract.UserDTO, error) {

	// Get Friendships
	friendships, err := u.repository.FriendshipRepository.GetFriendshipsBySearchIdAndState(searchId, state)

	if err != nil {
		return nil, err
	}

	if len(*friendships) == 0 {
		return &[]contract.UserDTO{}, nil
	}

	// Get Users
	var users []contract.UserDTO

	for _, friendship := range *friendships {

		// Get user by id
		user, err := u.repository.UserRepository.GetUserById(friendship.UserId)

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}
