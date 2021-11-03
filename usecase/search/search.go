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

func (usecase SearchUseCase) GetSearchesByUserSub(sub string) (*[]contract.SearchDTO, error) {

	user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	if err != nil {
		return nil, err
	}

	// Get user searches
	searches, err := usecase.repository.SearchRepository.GetSearchesByUserId(user.Id)

	if err != nil {
		return nil, err
	}


	return searches, err
}

func (usecase SearchUseCase) GetSearchesByFriendshipId(sub string) (*[]contract.SearchDTO, error) {
	return nil, nil
	// user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	// friendship, err := usecase.repository.FriendshipRepository.GetFriendshipsByUserId(user.Id)
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
