package model

import (
	"github.com/nightborn-be/invoice-backend/contract"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string
	LastName      string
	Email         string
	ExternalId    string
	Devices       []Device `gorm:"one2many:user_devices;"`
	Friendships   []Friendship `gorm:"one2many:user_friendships;"`
	Searches      []Search `gorm:"one2many:user_searches;"`
}

func ToUserDTO(user User) contract.UserDTO {
	return contract.UserDTO{
		Id:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		ExternalId: user.ExternalId,
	}
}

func ToUser(userDTO contract.UserDTO) User {
	return User{
		Model: gorm.Model{
			ID: userDTO.Id,
		},
		FirstName:  userDTO.FirstName,
		LastName:   userDTO.LastName,
		Email:      userDTO.Email,
		ExternalId: userDTO.ExternalId,
	}
}

func ToUserDTOs(Users []User) []contract.UserDTO {
	UserDTOs := make([]contract.UserDTO, len(Users))

	for i, item := range Users {
		UserDTOs[i] = ToUserDTO(item)
	}

	return UserDTOs
}

func ToUsers(UserDTOs []contract.UserDTO) []User {
	Users := make([]User, len(UserDTOs))

	for i, item := range UserDTOs {
		Users[i] = ToUser(item)
	}

	return Users
}
