package search_usecase

import (
	"errors"
	"log"
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

	searches, err := usecase.repository.SearchRepository.GetSearchesByUserId(user.Id)

	return searches, err
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
	log.Default().Println(searchDTO)
	search, err := usecase.repository.SearchRepository.AddSearch(searchDTO)
	return search, err
}

func (usecase SearchUseCase) ModifySearch(searchDTO contract.SearchDTO) (*contract.SearchDTO, error) {
	user, err := usecase.repository.SearchRepository.ModifySearch(searchDTO)
	return user, err
}
