package friendship_repository

import (

	"gorm.io/gorm"
)


type FriendshipRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) FriendshipRepository {
	return FriendshipRepository{
		database: db,
	}
}


