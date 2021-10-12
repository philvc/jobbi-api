package organisation_repository

import (
	"errors"

	contract "github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/database/model"
	"gorm.io/gorm"
)

type OrganisationRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) OrganisationRepository {
	return OrganisationRepository{
		database: db,
	}
}

func (repository OrganisationRepository) GetOrganisationsByUserId(userId uint) (*[]contract.OrganisationDTO, error) {
	var organisations []model.Organisation
	var user model.User

	if err := repository.database.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&user).Association("Organisations").Find(&organisations); err != nil {
		return nil, err
	}

	organisationDTOs := model.ToOrganisationDTOs(organisations)

	return &organisationDTOs, nil
}

func (repository OrganisationRepository) GetOrganisationById(organisationId string) (*contract.OrganisationDTO, error) {
	var organisation model.Organisation

	if err := repository.database.Where("id = ?", organisationId).First(&organisation).Error; err != nil {
		return nil, errors.New("no organisation")
	}

	organisationDTO := model.ToOrganisationDTO(organisation)

	return &organisationDTO, nil
}

func (repository OrganisationRepository) AddOrganisation(organisationDTO contract.OrganisationDTO) (*contract.OrganisationDTO, error) {

	organisation := model.ToOrganisation(organisationDTO)

	if err := repository.database.Create(&organisation).Error; err != nil {
		return nil, errors.New("failed to create organisation")
	}

	organisationDTO = model.ToOrganisationDTO(organisation)

	return &organisationDTO, nil
}

func (repository OrganisationRepository) ModifyOrganisation(organisationDTO contract.OrganisationDTO) (*contract.OrganisationDTO, error) {

	organisation := model.ToOrganisation(organisationDTO)

	repository.database.Model(&organisation).Where("id = ?", organisation.ID).Updates(map[string]interface{}{"name": organisation.Name,
		"slug": organisation.Slug, "address": organisation.Address, "vat_number": organisation.VATNumber, "owner": organisation.Owner,
		"swift": organisation.Swift, "bank_account": organisation.BankAccount, "company_color": organisation.CompanyColor,
		"contact_phone_number": organisation.ContactPhoneNumber, "contact_email": organisation.ContactEmail})

	organisationDTO = model.ToOrganisationDTO(organisation)

	return &organisationDTO, nil
}
