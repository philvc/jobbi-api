package organisation_usecase

import (
	"errors"
	"log"

	"github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/repository"
)

type OrganisationUsecase struct {
	repository repository.Repository
}

// Returns an instance of a organisation use-case
func Default(repository repository.Repository) OrganisationUsecase {
	return OrganisationUsecase{
		repository: repository,
	}
}

func (usecase OrganisationUsecase) GetOrganisationsByUserSub(sub string) (*[]contract.OrganisationDTO, error) {

	user, err := usecase.repository.UserRepository.GetUserBySub(sub)

	if err != nil {
		return nil, err
	}

	organisations, err := usecase.repository.OrganisationRepository.GetOrganisationsByUserId(user.Id)

	return organisations, err
}

func (usecase OrganisationUsecase) GetOrganisationById(organisationId string) (*contract.OrganisationDTO, error) {
	organisation, err := usecase.repository.OrganisationRepository.GetOrganisationById(organisationId)
	return organisation, err
}

func (usecase OrganisationUsecase) AddOrganisation(organisationDTO contract.OrganisationDTO) (*contract.OrganisationDTO, error) {

	if organisationDTO.Name == "" {
		return nil, errors.New("missing information")
	}

	if organisationDTO.VATNumber == "" {
		return nil, errors.New("missing information")
	}
	log.Default().Println(organisationDTO)
	organisation, err := usecase.repository.OrganisationRepository.AddOrganisation(organisationDTO)
	return organisation, err
}

func (usecase OrganisationUsecase) ModifyOrganisation(organisationDTO contract.OrganisationDTO) (*contract.OrganisationDTO, error) {
	user, err := usecase.repository.OrganisationRepository.ModifyOrganisation(organisationDTO)
	return user, err
}
