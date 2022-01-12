package model

import "github.com/philvc/jobbi-api/contract"

type Follower struct {
	Base
	UserID   string
	SearchID string
	Type     string
}


func ToFollowerDto(follower Follower) contract.FollowerDTO {
	return contract.FollowerDTO{
		Id: follower.ID,
		UserId: follower.UserID,
		SearchId: follower.SearchID,
	}
}

func ToFollower(followerDto contract.FollowerDTO) Follower {
	return Follower{
		Base: Base{
			ID: followerDto.Id,
		},
		SearchID: followerDto.SearchId,
		UserID: followerDto.UserId,
	}
}

func ToFollowerDTOs(followers []Follower) []contract.FollowerDTO {
	FollowerDTOs := make([]contract.FollowerDTO, len(followers))

	for i, item := range followers {
		FollowerDTOs[i] = ToFollowerDto(item)
	}

	return FollowerDTOs
}

func ToFollowers(followerDtos []contract.FollowerDTO) []Follower {
	Followers := make([]Follower, len(followerDtos))

	for i, item := range followerDtos {
		Followers[i] = ToFollower(item)
	}

	return Followers
}

