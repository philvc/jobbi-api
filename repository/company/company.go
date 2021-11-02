package company_repository

import (
	"errors"

	"github.com/google/uuid"
	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) CompanyRepository {
	return CompanyRepository{
		database: db,
	}
}

func (repository CompanyRepository) GetCompaniesBySearchId(searchId string) (*[]contract.CompanyDTO, error) {
	var Companies []model.Company
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&search).Association("Companies").Find(&Companies); err != nil {
		return nil, err
	}

	CompanyDTOs := model.ToCompanyDTOs(Companies)

	return &CompanyDTOs, nil
}

func (repository CompanyRepository) GetCompanyById(CompanyId string) (*contract.CompanyDTO, error) {
	var Company model.Company

	if err := repository.database.Where("id = ?", CompanyId).First(&Company).Error; err != nil {
		return nil, errors.New("no Company")
	}

	CompanyDTO := model.ToCompanyDTO(Company)

	return &CompanyDTO, nil
}

func (repository CompanyRepository) AddCompany(CompanyDTO contract.CompanyDTO) (*contract.CompanyDTO, error) {

	Company := model.ToCompany(CompanyDTO)

	id := uuid.New()

	Company.ID = id.String()

	if err := repository.database.Create(&Company).Error; err != nil {
		return nil, errors.New("failed to create Company")
	}

	CompanyDTO = model.ToCompanyDTO(Company)

	return &CompanyDTO, nil
}

func (repository CompanyRepository) ModifyCompany(CompanyDTO contract.CompanyDTO) (*contract.CompanyDTO, error) {

	Company := model.ToCompany(CompanyDTO)

	repository.database.Model(&Company).Where("id = ?", Company.ID).Updates(map[string]interface{}{
		"link": Company.Link, "description": Company.Description, "title": Company.Title})

	CompanyDTO = model.ToCompanyDTO(Company)

	return &CompanyDTO, nil
}

func (repository CompanyRepository) DeleteCompany(companyId string) (bool, error) {

	var Company model.Company

	if err := repository.database.Where("id = ?", companyId).First(&Company).Error; err != nil {
		return false, errors.New("no company")
	}

	if err := repository.database.Delete(&Company, companyId).Error; err != nil {
		return false, errors.New("delete offer failed")
	}

	return true, nil
}
