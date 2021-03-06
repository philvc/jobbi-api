package model

import (
	"github.com/philvc/jobbi-api/contract"
)

type Company struct {
	Base
	Description string
	Title       string
	Link 		string
	SearchID	string
	UserID 		string
}

func ToCompanyDTO(Company Company) contract.CompanyDTO {
	return contract.CompanyDTO{
		Id:          Company.Base.ID,
		Description: Company.Description,
		Title:       Company.Title,
		Link: 		Company.Link,
		SearchID: 	Company.SearchID,
		UserID: Company.UserID,
	}
}

func ToCompany(CompanyDTO contract.CompanyDTO) Company {
	return Company{
		Base: Base{
			ID: CompanyDTO.Id,
		},
		Link: CompanyDTO.Link,
		Description: CompanyDTO.Description,
		Title:       CompanyDTO.Title,
		SearchID: 	CompanyDTO.SearchID,
		UserID: CompanyDTO.UserID,
	}
}

func ToCompanyDTOs(Companies []Company) []contract.CompanyDTO {
	CompanyDTOs := make([]contract.CompanyDTO, len(Companies))

	for i, item := range Companies {
		CompanyDTOs[i] = ToCompanyDTO(item)
	}

	return CompanyDTOs
}

func ToCompanies(CompaniesDTO []contract.CompanyDTO) []Company {
	Companies := make([]Company, len(CompaniesDTO))

	for i, item := range CompaniesDTO {
		Companies[i] = ToCompany(item)
	}

	return Companies
}
