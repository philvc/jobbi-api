package search_usecase

import (
	"errors"

	constant "github.com/philvc/jobbi-api/constants"
	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type SearchUseCase struct {
	repository repository.Repository
}

// Returns an instance of a search use-case
func Default(repository repository.Repository) SearchUseCase {
	return SearchUseCase{
		repository: repository,
	}
}

// Get My search
func (usecase SearchUseCase) GetMySearch(sub string) (*contract.MySearchDTO, error) {

	user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	if err != nil {
		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetMySearch(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get posts by search id
func (usecase SearchUseCase) GetPostsBySearchId(searchId string) (*[]contract.PostDTOBySearchId, error) {

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetPostsBySearchId(searchId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get participants by search id
func (usecase SearchUseCase) GetParticipantsBySearchId(searchId string) (*[]contract.ParticipantDTOForSearchById, error) {

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetParticipantsBySearchId(searchId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get Shared searches
func (usecase SearchUseCase) GetSharedSearches(sub string) (*[]contract.SharedSearchDTO, error) {
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetSharedSearches(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get Followed searches
func (usecase SearchUseCase) GetFollowedSearches(sub string) (*[]contract.FollowedSearchDTO, error) {
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return nil, err
	}

	// Get user searches
	response, err := usecase.repository.SearchRepository.GetFollowedSearches(user.Id)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (usecase SearchUseCase) GetSearchById(searchId string) (*contract.SearchDTOById, error) {

	// Check access rights: owner or friend or follower or public

	search, err := usecase.repository.SearchRepository.GetSearchById(searchId)
	return search, err
}

func (usecase SearchUseCase) AddSearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	if searchDTO.Title == "" {
		return nil, errors.New(constant.ErrorMissingTitle)
	}

	if searchDTO.Description == "" {
		return nil, errors.New(constant.ErrorMissingDescription)
	}

	if searchDTO.Type == "" {
		return nil, errors.New(constant.ErrorMissingType)
	}

	// Check if user has already an existing search
	existingSearch, _ := usecase.repository.SearchRepository.GetMySearch(searchDTO.UserID)
	if existingSearch.Id != "" {
		return nil, errors.New(constant.ErrorAlreadyExistingSearch)
	}

	// Add search repository
	newSearch, err := usecase.repository.SearchRepository.AddSearch(searchDTO)

	return newSearch, err
}

func (usecase SearchUseCase) ModifySearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {
	search, err := usecase.repository.SearchRepository.ModifySearch(searchDTO)
	return search, err
}

func (usecase SearchUseCase) IsOwner(sub string, searchId string) bool {

	// Get user
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false
	}

	ok := usecase.repository.SearchRepository.IsSearchOwner(user.Id, searchId)

	return ok
}

func (usecase SearchUseCase) IsPublic(searchId string) bool {

	ok := usecase.repository.SearchRepository.IsPublic(searchId)

	return ok
}

func (usecase SearchUseCase) IsFriend(sub string, searchId string) bool {

	// Get user
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	if err != nil {
		return false
	}

	ok := usecase.repository.SearchRepository.IsFriend(user.Id, searchId)

	return ok
}
