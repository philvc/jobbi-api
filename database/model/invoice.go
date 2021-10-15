package model

import (
	"time"

	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Number      string
	Type        int
	Date        time.Time
	Description string
	Amount      float32
	SubTotal    float32
	VAT         float32
	Total       float32
	ClientId    int64
}

func ToInvoiceDTO(invoice Invoice) contract.InvoiceDTO {
	return contract.InvoiceDTO{
		Id:          invoice.ID,
		Number:      invoice.Number,
		Type:        invoice.Type,
		Date:        invoice.Date,
		Description: invoice.Description,
		Amount:      invoice.Amount,
		SubTotal:    invoice.SubTotal,
		VAT:         invoice.VAT,
		Total:       invoice.Total,
		ClientId:    invoice.ClientId,
	}
}

func ToInvoice(invoiceDTO contract.InvoiceDTO) Invoice {
	return Invoice{
		Model: gorm.Model{
			ID: invoiceDTO.Id,
		},
		Number:      invoiceDTO.Number,
		Type:        invoiceDTO.Type,
		Date:        invoiceDTO.Date,
		Description: invoiceDTO.Description,
		Amount:      invoiceDTO.Amount,
		SubTotal:    invoiceDTO.SubTotal,
		VAT:         invoiceDTO.VAT,
		Total:       invoiceDTO.Total,
		ClientId:    invoiceDTO.ClientId,
	}
}

func ToInvoiceDTOs(invoices []Invoice) []contract.InvoiceDTO {
	invoiceDTOs := make([]contract.InvoiceDTO, len(invoices))

	for i, item := range invoices {
		invoiceDTOs[i] = ToInvoiceDTO(item)
	}

	return invoiceDTOs
}

func ToInvoices(invoiceDTOs []contract.InvoiceDTO) []Invoice {
	invoices := make([]Invoice, len(invoiceDTOs))

	for i, item := range invoiceDTOs {
		invoices[i] = ToInvoice(item)
	}

	return invoices
}
