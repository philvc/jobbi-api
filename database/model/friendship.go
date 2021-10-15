package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	State     uint
	UserID uint
	SearchID uint
	Answers []Answer`gorm:"foreignKey:FriendshipID"`
}

func ToFriendshipDTO(friendship Friendship) contract.FriendshipDTO {
	return contract.FriendshipDTO{
		Id:         friendship.ID,
		State:  friendship.State,
	}
}

func ToFriendship(friendshipDTO contract.FriendshipDTO) Friendship {
	return Friendship{
		Model: gorm.Model{
			ID: friendshipDTO.Id,
		},
		State:  friendshipDTO.State,
	}
}

func ToFriendshipDTOs(friendships []Friendship) []contract.FriendshipDTO {
	FriendshipDTOs := make([]contract.FriendshipDTO, len(friendships))

	for i, item := range friendships {
		FriendshipDTOs[i] = ToFriendshipDTO(item)
	}

	return FriendshipDTOs
}

func ToFriendships(friendshipDTOs []contract.FriendshipDTO) []Friendship {
	Friendships := make([]Friendship, len(friendshipDTOs))

	for i, item := range friendshipDTOs {
		Friendships[i] = ToFriendship(item)
	}

	return Friendships
}
