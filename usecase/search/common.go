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

func (usecase SearchUseCase) hasSearchAccess(userId string, searchId string) (bool, error) {

	// Check if search is public
	isPublic := usecase.IsPublic(searchId)
	if !isPublic {

		// Check if user is owner
		IsSearchOwner := usecase.IsSearchOwner(userId, searchId)
		if !IsSearchOwner {

			// Check if user is friend or follower
			isFriend := usecase.IsFriend(userId, searchId)

			if !isFriend {
				return false, errors.New(constant.ErrorMissingAccess)
			}
		}
	}
	return true, nil
}
