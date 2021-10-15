package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Organisation struct {
	gorm.Model
	Name               string
	Slug               string
	Address            string
	VATNumber          string
	Owner              uint
	Swift              string
	BankAccount        string
	CompanyColor       string
	ContactPhoneNumber string
	ContactEmail       string
	Users              []User `gorm:"many2many:organisation_users;"`
}

func ToOrganisationDTO(organisation Organisation) contract.OrganisationDTO {
	return contract.OrganisationDTO{
		Id:                 organisation.ID,
		Name:               organisation.Name,
		Slug:               organisation.Slug,
		Address:            organisation.Address,
		VATNumber:          organisation.VATNumber,
		Owner:              organisation.Owner,
		Swift:              organisation.Swift,
		BankAccount:        organisation.BankAccount,
		CompanyColor:       organisation.CompanyColor,
		ContactPhoneNumber: organisation.ContactPhoneNumber,
		ContactEmail:       organisation.ContactEmail,
	}
}

func ToOrganisation(organisationDTO contract.OrganisationDTO) Organisation {
	return Organisation{
		Model: gorm.Model{
			ID: organisationDTO.Id,
		},
		Name:               organisationDTO.Name,
		Slug:               organisationDTO.Slug,
		Address:            organisationDTO.Address,
		VATNumber:          organisationDTO.VATNumber,
		Owner:              organisationDTO.Owner,
		Swift:              organisationDTO.Swift,
		BankAccount:        organisationDTO.BankAccount,
		CompanyColor:       organisationDTO.CompanyColor,
		ContactPhoneNumber: organisationDTO.ContactPhoneNumber,
		ContactEmail:       organisationDTO.ContactEmail,
	}
}

func ToOrganisationDTOs(organisations []Organisation) []contract.OrganisationDTO {
	organisationDTOs := make([]contract.OrganisationDTO, len(organisations))

	for i, item := range organisations {
		organisationDTOs[i] = ToOrganisationDTO(item)
	}

	return organisationDTOs
}

func ToOrganisations(organisationDTOs []contract.OrganisationDTO) []Organisation {
	organisations := make([]Organisation, len(organisationDTOs))

	for i, item := range organisationDTOs {
		organisations[i] = ToOrganisation(item)
	}

	return organisations
}
