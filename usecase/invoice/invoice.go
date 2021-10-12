package invoice_usecase

import (
	"errors"

	contract "github.com/nightborn-be/invoice-backend/contract"
	"github.com/nightborn-be/invoice-backend/repository"
)

type InvoiceUsecase struct {
	repository repository.Repository
}

// Returns an instance of a invoice use-case
func Default(repository repository.Repository) InvoiceUsecase {
	return InvoiceUsecase{
		repository: repository,
	}
}

func (usecase InvoiceUsecase) GetInvoicesByClientId(organisationId string) (*[]contract.InvoiceDTO, error) {
	invoices, err := usecase.repository.InvoiceRepository.GetInvoicesByClientId(organisationId)
	return invoices, err
}

func (usecase InvoiceUsecase) GetInvoicesByOrganisationId(organisationId string) (*[]contract.InvoiceDTO, error) {
	invoices, err := usecase.repository.InvoiceRepository.GetInvoicesByOrganisationId(organisationId)
	return invoices, err
}

func (usecase InvoiceUsecase) GetInvoiceById(invoiceId string) (*contract.InvoiceDTO, error) {
	invoice, err := usecase.repository.InvoiceRepository.GetInvoiceById(invoiceId)
	return invoice, err
}

func (usecase InvoiceUsecase) AddInvoice(invoiceDTO contract.InvoiceDTO) (*contract.InvoiceDTO, error) {

	if invoiceDTO.Number == "" {
		return nil, errors.New("missing information")
	}

	if invoiceDTO.Description == "" {
		return nil, errors.New("missing information")
	}

	if invoiceDTO.Amount < 0 {
		return nil, errors.New("negative value")
	}

	if invoiceDTO.SubTotal < 0 {
		return nil, errors.New("negative value")
	}

	if invoiceDTO.VAT < 0 {
		return nil, errors.New("negative value")
	}

	if invoiceDTO.Total < 0 {
		return nil, errors.New("negative value")
	}

	if invoiceDTO.ClientId < 0 {
		return nil, errors.New("missing information")
	}

	// TODO : Should verify that a invoice number is unique

	// Should verify that invoice date is later or equal to last date

	invoice, err := usecase.repository.InvoiceRepository.AddInvoice(invoiceDTO)
	return invoice, err
}

func (usecase InvoiceUsecase) ModifyInvoice(invoiceDTO contract.InvoiceDTO) (*contract.InvoiceDTO, error) {
	invoice, err := usecase.repository.InvoiceRepository.ModifyInvoice(invoiceDTO)
	return invoice, err
}
