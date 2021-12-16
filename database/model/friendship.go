package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type Friendship struct {
	Base
	State    uint
	UserID   string
	SearchID string
}

func ToFriendshipDTO(friendship Friendship) contract.FriendshipDTO {
	return contract.FriendshipDTO{
		Id:       friendship.Base.ID,
		State:    friendship.State,
		UserId:   friendship.UserID,
		SearchId: friendship.SearchID,
	}
}

func ToFriendship(friendshipDTO contract.FriendshipDTO) Friendship {
	return Friendship{
		Base: Base{
			ID: friendshipDTO.Id,
		},
		State:    friendshipDTO.State,
		UserID:   friendshipDTO.UserId,
		SearchID: friendshipDTO.SearchId,
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
