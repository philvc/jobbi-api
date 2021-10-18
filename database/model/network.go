package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Network struct {
	gorm.Model
	Description string
	LastName       string
	FirstName       string
	Email       string
	PhoneNumber       string
	Link 		string
	SearchID		uint
	UserID uint
}

func ToNetworkDTO(Network Network) contract.NetworkDTO {
	return contract.NetworkDTO{
		Id:          Network.ID,
		Description: Network.Description,
		FirstName:       Network.FirstName,
		LastName:       Network.LastName,
		PhoneNumber:       Network.PhoneNumber,
		Email:       Network.Email,
		Link: 		Network.Link,
		SearchID: 	Network.SearchID,
		UserID: Network.UserID,
	}
}

func ToNetwork(NetworkDTO contract.NetworkDTO) Network {
	return Network{
		Model: gorm.Model{
			ID: NetworkDTO.Id,
		},
		Link: NetworkDTO.Link,
		Description: NetworkDTO.Description,
		LastName:       NetworkDTO.LastName,
		FirstName:       NetworkDTO.FirstName,
		PhoneNumber:       NetworkDTO.PhoneNumber,
		Email:       NetworkDTO.Email,
		SearchID: 	NetworkDTO.SearchID,
		UserID: NetworkDTO.UserID,
	}
}

func ToNetworkDTOs(Networks []Network) []contract.NetworkDTO {
	NetworkDTOs := make([]contract.NetworkDTO, len(Networks))

	for i, item := range Networks {
		NetworkDTOs[i] = ToNetworkDTO(item)
	}

	return NetworkDTOs
}

func ToNetworks(NetworksDTO []contract.NetworkDTO) []Network {
	Networks := make([]Network, len(NetworksDTO))

	for i, item := range NetworksDTO {
		Networks[i] = ToNetwork(item)
	}

	return Networks
}