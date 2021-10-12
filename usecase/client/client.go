package client_usecase

import (
	"errors"

	contract "github.com/nightborn-be/invoice-backend/contract"
	repository "github.com/nightborn-be/invoice-backend/repository"
)

type ClientUsecase struct {
	repository repository.Repository
}

// Returns an instance of a client use-case
func Default(repository repository.Repository) ClientUsecase {
	return ClientUsecase{
		repository: repository,
	}
}

func (usecase ClientUsecase) GetClientsByOrganisationId(organisationId string) (*[]contract.ClientDTO, error) {
	clients, err := usecase.repository.ClientRepository.GetClientsByOrganisationId(organisationId)
	return clients, err
}

func (usecase ClientUsecase) GetClientById(clientId int64) (*contract.ClientDTO, error) {
	client, err := usecase.repository.ClientRepository.GetClientById(clientId)
	return client, err
}

func (usecase ClientUsecase) AddClient(clientDTO contract.ClientDTO) (*contract.ClientDTO, error) {

	if clientDTO.Name == "" {
		return nil, errors.New("missing information")
	}

	if clientDTO.Email == "" {
		return nil, errors.New("missing information")
	}

	if clientDTO.Address == "" {
		return nil, errors.New("missing information")
	}

	client, err := usecase.repository.ClientRepository.AddClient(clientDTO)
	return client, err
}

func (usecase ClientUsecase) ModifyClient(clientDTO contract.ClientDTO) (*contract.ClientDTO, error) {
	client, err := usecase.repository.ClientRepository.ModifyClient(clientDTO)
	return client, err
}
