package search_usecase

import (
	"errors"

	constant "github.com/philvc/jobbi-api/constants"
	"github.com/philvc/jobbi-api/contract"
)

func (usecase SearchUseCase) IsPostOwner(userId string, postId string) bool {

	ok := usecase.repository.SearchRepository.IsPostOwner(userId, postId)

	return ok
}

func (usecase SearchUseCase) IsSearchOwner(userId string, searchId string) bool {

	ok := usecase.repository.SearchRepository.IsSearchOwner(userId, searchId)

	return ok
}

func (usecase SearchUseCase) IsPublic(searchId string) bool {

	ok := usecase.repository.SearchRepository.IsPublic(searchId)

	return ok
}

func (usecase SearchUseCase) IsFriend(userId string, searchId string) bool {

	ok := usecase.repository.SearchRepository.IsFriend(userId, searchId)

	return ok
}

func (usecase SearchUseCase) IsSearchExist(searchId string) (*contract.SearchDTO, error) {
	search, err := usecase.repository.SearchRepository.IsSearchExist(searchId)
	if err != nil {
		return nil, err
	}

	return search, nil
}

func (usecase SearchUseCase) IsPostExist(postId string) (*contract.PostDTO, error) {
	post, err := usecase.repository.SearchRepository.IsPostExist(postId)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (usecase SearchUseCase) IsPostExistForSearch(postId string, searchId string) (*contract.PostDTO, error) {
	post, err := usecase.repository.SearchRepository.IsPostExistForSearch(postId, searchId)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (usecase SearchUseCase) hasSearchAccess(userId string, searchId string) (bool, error) {

	// Check if search is public
	isPublic := usecase.IsPublic(searchId)
	if !isPublic {

		// Check if user is owner
		IsSearchOwner := usecase.IsSearchOwner(userId, searchId)
		if !IsSearchOwner {

			// Check if user is friend
			isFriend := usecase.IsFriend(userId, searchId)

			if !isFriend {
				return false, errors.New(constant.ErrorMissingAccess)
			}
		}
	}
	return true, nil
}

func (usecase SearchUseCase) IsFollowerExist(searchId string, userId string) (*contract.FollowerDTO, error) {
	follower, err := usecase.repository.SearchRepository.IsFollowerExist(searchId, userId)
	if err != nil {
		return nil, err
	}

	return follower, nil
}

func (usecase SearchUseCase) IsFriendshipExist(searchId string, userId string) (*contract.FriendshipDTO, error) {
	friendship, err := usecase.repository.SearchRepository.IsFriendshipExist(searchId, userId)
	if err != nil {
		return nil, err
	}

	return friendship, nil
}

func (usecase SearchUseCase) IsFollowerExistById(followerId string) (*contract.FollowerDTO, error) {
	// Call repo
	follower, err := usecase.repository.SearchRepository.GetFollowerById(followerId)
	if err != nil {
		return nil, err
	}

	return follower, nil
}
