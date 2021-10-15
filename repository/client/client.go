package client_repository

import (
	"errors"
	"fmt"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type ClientRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) ClientRepository {
	return ClientRepository{
		database: db,
	}
}

func (repository ClientRepository) GetClientsByOrganisationId(organisationId string) (*[]contract.ClientDTO, error) {
	var clients []model.Client

	if err := repository.database.Where("organisation_id = ?", organisationId).Find(&clients).Error; err != nil {
		return nil, errors.New("no clients")
	}

	clientDTOs := model.ToClientDTOs(clients)

	return &clientDTOs, nil
}

func (repository ClientRepository) GetClientById(clientId int64) (*contract.ClientDTO, error) {
	var client model.Client

	if err := repository.database.Where("id = ?", fmt.Sprint(clientId)).First(&client).Error; err != nil {
		return nil, errors.New("no client")
	}

	clientDTO := model.ToClientDTO(client)

	return &clientDTO, nil
}

func (repository ClientRepository) AddClient(clientDTO contract.ClientDTO) (*contract.ClientDTO, error) {

	client := model.ToClient(clientDTO)

	if err := repository.database.Create(&client).Error; err != nil {
		return nil, errors.New("failed to create client")
	}

	clientDTO = model.ToClientDTO(client)

	return &clientDTO, nil
}

func (repository ClientRepository) ModifyClient(clientDTO contract.ClientDTO) (*contract.ClientDTO, error) {

	client := model.ToClient(clientDTO)

	repository.database.Model(&client).Where("id = ?", client.ID).Updates(map[string]interface{}{"name": client.Name,
		"email": client.Email, "address": client.Address, "vat_number": client.VATNumber,
		"vat_included": client.VATIncluded})

	clientDTO = model.ToClientDTO(client)

	return &clientDTO, nil
}
