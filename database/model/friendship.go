package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	State     uint
	FirstName string
	LastName string
	Email string
	UserID uint
	SearchID uint
	Answers []Answer`gorm:"foreignKey:FriendshipID"`
}

func ToFriendshipDTO(friendship Friendship) contract.FriendshipDTO {
	return contract.FriendshipDTO{
		Id:         friendship.ID,
		State:  friendship.State,
		FirstName:       friendship.FirstName,
		LastName:       friendship.LastName,
		Email:       friendship.Email,
	}
}

func ToFriendship(friendshipDTO contract.FriendshipDTO) Friendship {
	return Friendship{
		Model: gorm.Model{
			ID: friendshipDTO.Id,
		},
		State:  friendshipDTO.State,
		FirstName:       friendshipDTO.FirstName,
		LastName:       friendshipDTO.LastName,
		Email:       friendshipDTO.Email,
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
