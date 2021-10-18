package company_usecase

import (
	"errors"

	"github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/repository"
)

type CompanyUseCase struct {
	repository repository.Repository
}

// Returns an instance of a Company use-case
func Default(repository repository.Repository) CompanyUseCase {
	return CompanyUseCase{
		repository: repository,
	}
}


func (usecase CompanyUseCase) GetCompaniesBySearchId(searchId string) (*[]contract.CompanyDTO, error) {
	Companies, err := usecase.repository.CompanyRepository.GetCompaniesBySearchId(searchId)
	return Companies, err
}

func (usecase CompanyUseCase) GetCompanyById(CompanyId string) (*contract.CompanyDTO, error) {
	Company, err := usecase.repository.CompanyRepository.GetCompanyById(CompanyId)
	return Company, err
}

func (usecase CompanyUseCase) AddCompany(CompanyDTO contract.CompanyDTO) (*contract.CompanyDTO, error) {

	if CompanyDTO.Title == "" {
		return nil, errors.New("missing title")
	}

	if CompanyDTO.Description == "" {
		return nil, errors.New("missing description")
	}
	if CompanyDTO.Link == "" {
		return nil, errors.New("missing link")
	}
	Company, err := usecase.repository.CompanyRepository.AddCompany(CompanyDTO)
	return Company, err
}

func (usecase CompanyUseCase) ModifyCompany(CompanyDTO contract.CompanyDTO) (*contract.CompanyDTO, error) {
	user, err := usecase.repository.CompanyRepository.ModifyCompany(CompanyDTO)
	return user, err
}
