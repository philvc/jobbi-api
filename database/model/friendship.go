package model

import (
	"github.com/nightborn-be/invoice-backend/contract"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	State     uint
	User 	User`gorm:"many2one:friendship_user;"`
	Answers []Answer`gorm:"one2many:friendship_answers;"`
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
