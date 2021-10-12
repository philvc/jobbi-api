package model

import (
	"github.com/nightborn-be/invoice-backend/contract"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Token string
}

func ToDeviceDTO(Device Device) contract.DeviceDTO {
	return contract.DeviceDTO{
		Id:         Device.ID,
		Token:  Device.Token,
	}
}

func ToDevice(DeviceDTO contract.DeviceDTO) Device {
	return Device{
		Model: gorm.Model{
			ID: DeviceDTO.Id,
		},
		Token:  DeviceDTO.Token,
	}
}

