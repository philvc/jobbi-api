package network_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type NetworkUseCase struct {
	repository repository.Repository
}

// Returns an instance of a Network use-case
func Default(repository repository.Repository) NetworkUseCase {
	return NetworkUseCase{
		repository: repository,
	}
}


func (usecase NetworkUseCase) GetNetworksBySearchId(searchId string) (*[]contract.NetworkDTO, error) {
	Networks, err := usecase.repository.NetworkRepository.GetNetworksBySearchId(searchId)
	return Networks, err
}

func (usecase NetworkUseCase) GetNetworkById(NetworkId string) (*contract.NetworkDTO, error) {
	Network, err := usecase.repository.NetworkRepository.GetNetworkById(NetworkId)
	return Network, err
}

func (usecase NetworkUseCase) AddNetwork(NetworkDTO contract.NetworkDTO) (*contract.NetworkDTO, error) {

	if NetworkDTO.FirstName == "" {
		return nil, errors.New("missing firstName")
	}

	if NetworkDTO.Description == "" {
		return nil, errors.New("missing description")
	}

	Network, err := usecase.repository.NetworkRepository.AddNetwork(NetworkDTO)
	return Network, err
}

func (usecase NetworkUseCase) ModifyNetwork(NetworkDTO contract.NetworkDTO) (*contract.NetworkDTO, error) {
	user, err := usecase.repository.NetworkRepository.ModifyNetwork(NetworkDTO)
	return user, err
}
