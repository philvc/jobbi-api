package search_usecase

import (
	"errors"

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

// Get Shared searches
func (usecase SearchUseCase) GetSharedSearches(sub string)(*[]contract.SharedSearchDTO, error){
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
func (usecase SearchUseCase) GetFollowedSearches(sub string)(*[]contract.FollowedSearchDTO, error){
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


func (usecase SearchUseCase) GetSearchById(searchId string) (*contract.SearchDTO, error) {
	search, err := usecase.repository.SearchRepository.GetSearchById(searchId)
	return search, err
}

func (usecase SearchUseCase) AddSearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {

	if searchDTO.Title == "" {
		return nil, errors.New("missing title")
	}

	if searchDTO.Description == "" {
		return nil, errors.New("missing description")
	}
	search, err := usecase.repository.SearchRepository.AddSearch(searchDTO)
	return search, err
}

func (usecase SearchUseCase) ModifySearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {
	user, err := usecase.repository.SearchRepository.ModifySearch(searchDTO)
	return user, err
}
