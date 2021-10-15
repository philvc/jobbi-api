package invoice_repository

import (
	"errors"
	"log"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) InvoiceRepository {
	return InvoiceRepository{
		database: db,
	}
}

func (repository InvoiceRepository) GetInvoicesByOrganisationId(organisationId string) (*[]contract.InvoiceDTO, error) {
	var invoices []model.Invoice

	log.Default().Println(organisationId)
	if err := repository.database.Model(&invoices).Joins("left join clients on clients.id = invoices.client_id").Where("clients.organisation_id = ?", organisationId).Scan(&invoices).Error; err != nil {
		return nil, errors.New("no invoices for given organisation")
	}
	log.Default().Println(invoices)

	invoiceDTOs := model.ToInvoiceDTOs(invoices)

	return &invoiceDTOs, nil
}

func (repository InvoiceRepository) GetInvoicesByClientId(clientId string) (*[]contract.InvoiceDTO, error) {
	var invoices []model.Invoice

	if err := repository.database.Where("client_id = ?", clientId).Find(&invoices).Error; err != nil {
		return nil, errors.New("no invoices for given organisation")
	}

	invoiceDTOs := model.ToInvoiceDTOs(invoices)

	return &invoiceDTOs, nil
}

func (repository InvoiceRepository) GetInvoiceById(invoiceId string) (*contract.InvoiceDTO, error) {
	var invoice model.Invoice

	if err := repository.database.Where("id = ?", invoiceId).First(&invoice).Error; err != nil {
		return nil, errors.New("no invoice")
	}

	invoiceDTO := model.ToInvoiceDTO(invoice)

	return &invoiceDTO, nil
}

func (repository InvoiceRepository) AddInvoice(invoiceDTO contract.InvoiceDTO) (*contract.InvoiceDTO, error) {

	invoice := model.ToInvoice(invoiceDTO)

	if err := repository.database.Create(&invoice).Error; err != nil {
		return nil, errors.New("failed to create invoice")
	}

	invoiceDTO = model.ToInvoiceDTO(invoice)

	return &invoiceDTO, nil
}

func (repository InvoiceRepository) ModifyInvoice(invoiceDTO contract.InvoiceDTO) (*contract.InvoiceDTO, error) {

	invoice := model.ToInvoice(invoiceDTO)

	if err := repository.database.Model(&invoice).Updates(map[string]interface{}{"Number": invoice.Number,
		"Type": invoice.Type, "Date": invoice.Date, "Description": invoice.Description, "Amount": invoice.Amount,
		"SubTotal": invoice.SubTotal, "VAT": invoice.VAT, "Total": invoice.Total}); err != nil {
		return nil, errors.New("failed to update invoice")
	}

	invoiceDTO = model.ToInvoiceDTO(invoice)

	return &invoiceDTO, nil
}
