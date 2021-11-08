package friendship_usecase

import (
	"log"

	contract "github.com/philvc/jobbi-api/contract/search"
)

func (u *FriendshipUsecaseImplementation) GetFriendshipsByUserIdAndState(searchId string, state uint)(*contract.SearchWithOwnerAndFriends, error) {
	log.Print("get friendships by user id and state")
}
