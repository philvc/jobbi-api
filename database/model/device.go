package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type Device struct {
	Base
	Token string
	UserID uint
}

func ToDeviceDTO(Device Device) contract.DeviceDTO {
	return contract.DeviceDTO{
		Id:         Device.Base.ID,
		Token:  Device.Token,
	}
}

func ToDevice(DeviceDTO contract.DeviceDTO) Device {
	return Device{
		Base: Base{
			ID: DeviceDTO.Id,
		},
		Token:  DeviceDTO.Token,
	}
}

