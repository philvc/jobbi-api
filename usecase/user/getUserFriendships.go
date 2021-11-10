package user_usecase

import (

	"github.com/philvc/jobbi-api/contract"
)




func (u UserUsecase) GetUserFriendships(userId string, state int) (*[]contract.SearchWithOwnerAndFriendsDTO, error) {

	// Get friendships by user
	friendships, err := u.repository.FriendshipRepository.GetFriendshipsByUserIdAndState(userId, state)

	if len(*friendships) == 0 {
		return &[]contract.SearchWithOwnerAndFriendsDTO{}, nil
	}

	// Error
	if err != nil {
		return nil, err
	}

	// Get Searches
	var searches []contract.SearchDTO

	for _, friendship := range *friendships {
		// Get Search by id
		search, err := u.repository.SearchRepository.GetSearchById(friendship.SearchId)

		// Error
		if err != nil {
			return nil , err
		}

		searches = append(searches, *search)
	}

	// Get Owners

	var owners []contract.UserDTO


	for _, search := range searches {

		// Get search owner by id
		user, err := u.repository.UserRepository.GetUserById(search.UserID)

		// Error
		if err != nil {
			return nil, err
		}

		owners = append(owners, *user)
	}

	// Build response
	var results []contract.SearchWithOwnerAndFriendsDTO

	for index, search := range searches {

		// Result
		var result contract.SearchWithOwnerAndFriendsDTO

		// Friends for search 
		friends := []contract.UserDTO{}

		// Get Friendships by searchId
		friendships, err := u.repository.FriendshipRepository.GetFriendshipsBySearchIdAndState(search.Id, state)

		// Error
		if err != nil {
			return nil, err
		}
		
		// Get user and add to friends
		for _, friendship := range *friendships {
			user, err := u.repository.UserRepository.GetUserById(friendship.UserId)

			if err != nil {
				return nil, err
			}

			friends = append(friends, *user)
		}

		result.Owner = owners[index]
		result.Search = search
		result.Friends = friends

		results = append(results, result)
	}

	return &results, nil
}
