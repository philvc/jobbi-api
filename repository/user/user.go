package user_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) UserRepository {
	return UserRepository{
		database: db,
	}
}

func (repository UserRepository) GetUserBySub(sub string) (*contract.UserDTO, error) {
	var user model.User

	if err := repository.database.Where("external_id = ?", sub).First(&user).Error; err != nil {
		return nil, errors.New("no user")
	}

	userDTO := model.ToUserDTO(user)

	return &userDTO, nil
}

func (repository UserRepository) GetUserByEmail(email string) (*contract.UserDTO, error){
	var user model.User

	if err := repository.database.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("no user")
	}

	userDTO := model.ToUserDTO(user)

	return &userDTO, nil
}

func (repository UserRepository) AddUser(userDTO contract.UserDTO) (*contract.UserDTO, error) {

	user := model.ToUser(userDTO)

	if err := repository.database.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	userDTO = model.ToUserDTO(user)

	return &userDTO, nil
}

func (repository UserRepository) ModifyUser(userDTO contract.UserDTO) (*contract.UserDTO, error) {

	user := model.ToUser(userDTO)

	repository.database.Model(&user).Where("id = ?", user.ID).Updates(map[string]interface{}{"first_name": user.FirstName,
		"email": user.Email, "last_name": user.LastName})

	userDTO = model.ToUserDTO(user)

	return &userDTO, nil
}
