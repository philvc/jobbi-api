package friendship_usecase

import (
	"github.com/philvc/jobbi-api/contract/search"
	"github.com/philvc/jobbi-api/repository/friendship"
)

type FriendshipUsecase interface {
	GetFriendshipsByUserIdAndState(userId string, state uint)(*contract.SearchWithOwnerAndFriends, error)
}

type FriendshipUsecaseImplementation struct {
	repository friendship_repository.FriendshipRepositoryImplementation
}

// Returns an instance of a Friendship use-case
func Default(repository friendship_repository.FriendshipRepositoryImplementation) FriendshipUsecaseImplementation {
	return FriendshipUsecaseImplementation{
		repository: repository,
	}
}


