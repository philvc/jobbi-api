package user_repository

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
)

func (repository UserRepository) GetUserById(id string) (*contract.UserDTO, error) {
	var user model.User

	if err := repository.database.Where("id = ?", id).First(&user); err.Error != nil {
		return nil, errors.New("user-by-id-error")
	}

	userDto := model.ToUserDTO(user)

	return &userDto, nil
}
