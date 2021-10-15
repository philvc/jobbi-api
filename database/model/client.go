package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name           string
	Email          string
	Address        string
	VATNumber      string
	VATIncluded    bool
	OrganisationId uint
}

func ToClientDTO(client Client) contract.ClientDTO {
	return contract.ClientDTO{
		Id:             client.ID,
		Name:           client.Name,
		Email:          client.Email,
		Address:        client.Address,
		VATNumber:      client.VATNumber,
		VATIncluded:    client.VATIncluded,
		OrganisationId: client.OrganisationId,
	}
}

func ToClient(clientDTO contract.ClientDTO) Client {
	return Client{
		Model: gorm.Model{
			ID: clientDTO.Id,
		},
		Name:           clientDTO.Name,
		Email:          clientDTO.Email,
		Address:        clientDTO.Address,
		VATNumber:      clientDTO.VATNumber,
		VATIncluded:    clientDTO.VATIncluded,
		OrganisationId: clientDTO.OrganisationId,
	}
}

func ToClientDTOs(clients []Client) []contract.ClientDTO {
	clientDTOs := make([]contract.ClientDTO, len(clients))

	for i, item := range clients {
		clientDTOs[i] = ToClientDTO(item)
	}

	return clientDTOs
}

func ToClients(clientDTOs []contract.ClientDTO) []Client {
	clients := make([]Client, len(clientDTOs))

	for i, item := range clientDTOs {
		clients[i] = ToClient(item)
	}

	return clients
}
