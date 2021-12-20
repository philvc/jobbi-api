package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type User struct {
	Base
	FirstName   string
	LastName    string
	Email       string
	ExternalId  string
	AvatarUrl   string
	Searches    []Search     `gorm:"foreignKey:UserID"`
	Devices     []Device     `gorm:"foreignKey:UserID"`
	Friendships []Friendship `gorm:"foreignKey:UserID"`
	Offers      []Offer      `gorm:"foreignKey:UserID"`
	Companies   []Company    `gorm:"foreignKey:UserID"`
	Networks    []Network    `gorm:"foreignKey:UserID"`
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
		Base: Base{
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
