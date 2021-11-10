package friendship_usecase

import (
	"github.com/philvc/jobbi-api/repository"
)

type FriendshipUsecase struct {
	repository repository.Repository
}

// Returns an instance of a Friendship use-case
func Default(repository repository.Repository ) FriendshipUsecase {
	return FriendshipUsecase{
		repository: repository,
	}
}

