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

func (usecase FriendshipUseCase) GetFriendshipsBySearchId(searchId string, status uint) (*[]contract.FriendshipDTO, error) {

	Friendships, err := usecase.repository.FriendshipRepository.GetFriendshipsBySearchId(searchId, status)
	return Friendships, err
}

func (usecase FriendshipUseCase) GetFriendshipsBySub(sub string, status uint) (*[]contract.FriendshipDTO, error) {

	user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	if err != nil {
		return nil, err
	}

	Friendships, err := usecase.repository.FriendshipRepository.GetFriendshipsByUserId(user.Id)

	// filter friendships by status
	 filteredFriendships := make([]contract.FriendshipDTO, 0, len(*Friendships));

	 for _, friendship := range *Friendships {
		 if friendship.State == status {
			 filteredFriendships = append(filteredFriendships, friendship)
		 }
	 }
	

	return &filteredFriendships, err
}

func (usecase FriendshipUseCase) AddFriendship(FriendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {

	if FriendshipDTO.Email == "" {
		return nil, errors.New("missing email")
	}

	Friendship, err := usecase.repository.FriendshipRepository.AddFriendship(FriendshipDTO)
	return Friendship, err
}

func (usecase FriendshipUseCase) ModifyFriendship(FriendshipDTO contract.FriendshipDTO) (*contract.FriendshipDTO, error) {
	user, err := usecase.repository.FriendshipRepository.ModifyFriendship(FriendshipDTO)
	return user, err
}
