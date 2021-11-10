package user_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)



type UserUsecase struct {
	repository repository.Repository
}

// Returns an instance of a user use-case
func Default(repository repository.Repository) UserUsecase {
	return UserUsecase{
		repository: repository,
	}
}

func (usecase UserUsecase) GetUserBySub(sub string) (*contract.UserDTO, error) {
	user, err := usecase.repository.UserRepository.GetUserBySub(sub)
	return user, err
}

func (usecase UserUsecase) GetUserByEmail(email string) (*contract.UserDTO, error) {
	user, err := usecase.repository.UserRepository.GetUserByEmail(email)
	return user, err
}

func (usecase UserUsecase) AddUser(userDTO contract.UserDTO) (*contract.UserDTO, error) {

	// Verify if email already exist:
	// Get user by email
	existingUser, _ := usecase.repository.UserRepository.GetUserByEmail(userDTO.Email)

	if existingUser != nil {
		return nil, errors.New(("wrong account information"))
	}

	if userDTO.ExternalId == "" {
		return nil, errors.New("missing information")
	}

	if userDTO.Email == "" {
		return nil, errors.New("missing information")
	}

	user, err := usecase.repository.UserRepository.AddUser(userDTO)

	return user, err
}

func (usecase UserUsecase) ModifyUser(userDTO contract.UserDTO) (*contract.UserDTO, error) {

	if userDTO.Email == "" {
		return nil, errors.New("missing information")
	}

	user, err := usecase.repository.UserRepository.ModifyUser(userDTO)
	return user, err
}
