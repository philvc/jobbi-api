package friendship_repository

import (

	contract "github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type FriendshipRepository interface {
	GetFriendshipsByUserIdAndState(userId string, state uint) (*[]contract.FriendshipDTO, error)
}

type FriendshipRepositoryImplementation struct {
	database *gorm.DB
}

func Default(db *gorm.DB) FriendshipRepositoryImplementation {
	return FriendshipRepositoryImplementation{
		database: db,
	}
}


