package network_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type NetworkRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) NetworkRepository {
	return NetworkRepository{
		database: db,
	}
}

func (repository NetworkRepository) GetNetworksBySearchId(searchId string) (*[]contract.NetworkDTO, error) {
	var Networks []model.Network
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&search).Association("Networks").Find(&Networks); err != nil {
		return nil, err
	}

	NetworkDTOs := model.ToNetworkDTOs(Networks)

	return &NetworkDTOs, nil
}

func (repository NetworkRepository) GetNetworkById(NetworkId string) (*contract.NetworkDTO, error) {
	var Network model.Network

	if err := repository.database.Where("id = ?", NetworkId).First(&Network).Error; err != nil {
		return nil, errors.New("no Network")
	}

	NetworkDTO := model.ToNetworkDTO(Network)

	return &NetworkDTO, nil
}

func (repository NetworkRepository) AddNetwork(NetworkDTO contract.NetworkDTO) (*contract.NetworkDTO, error) {

	Network := model.ToNetwork(NetworkDTO)

	if err := repository.database.Create(&Network).Error; err != nil {
		return nil, errors.New("failed to create Network")
	}

	NetworkDTO = model.ToNetworkDTO(Network)

	return &NetworkDTO, nil
}

func (repository NetworkRepository) ModifyNetwork(NetworkDTO contract.NetworkDTO) (*contract.NetworkDTO, error) {

	Network := model.ToNetwork(NetworkDTO)

	repository.database.Model(&Network).Where("id = ?", Network.ID).Updates(map[string]interface{}{
		"link": Network.Link, "description": Network.Description, "first_name": Network.FirstName, "last_name": Network.LastName, "email": Network.Email, "phone_number": Network.PhoneNumber})
	
	NetworkDTO = model.ToNetworkDTO(Network)

	return &NetworkDTO, nil
}



func (repository NetworkRepository) DeleteNetwork(networkId string) (bool, error) {

	var Network model.Network

	if err := repository.database.Where("id = ?", networkId).First(&Network).Error; err != nil {
		return false, errors.New("no Network")
	}

	if err := repository.database.Delete(&Network, networkId).Error; err != nil {
		return false, errors.New("delete network failed")
	}

	return true, nil
}
