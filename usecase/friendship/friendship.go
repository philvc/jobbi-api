package friendship_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type FriendshipUseCase struct {
	repository repository.Repository
}

// Returns an instance of a Friendship use-case
func Default(repository repository.Repository) FriendshipUseCase {
	return FriendshipUseCase{
		repository: repository,
	}
}


func (usecase FriendshipUseCase) GetFriendshipsBySearchId(searchId string) (*[]contract.FriendshipDTO, error) {
	Friendships, err := usecase.repository.FriendshipRepository.GetFriendshipsBySearchId(searchId)
	return Friendships, err
}

func (usecase FriendshipUseCase) GetFriendshipById(FriendshipId string) (*contract.FriendshipDTO, error) {
	Friendship, err := usecase.repository.FriendshipRepository.GetFriendshipById(FriendshipId)
	return Friendship, err
}

func (usecase FriendshipUseCase) AddFriendship(FriendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	if FriendshipDTO.FirstName == "" {
		return nil, errors.New("missing firstName")
	}

	if FriendshipDTO.Description == "" {
		return nil, errors.New("missing description")
	}

	Friendship, err := usecase.repository.FriendshipRepository.AddFriendship(FriendshipDTO)
	return Friendship, err
}

func (usecase FriendshipUseCase) ModifyFriendship(FriendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {
	user, err := usecase.repository.FriendshipRepository.ModifyFriendship(FriendshipDTO)
	return user, err
}
