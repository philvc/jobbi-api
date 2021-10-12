package contract

import "time"

// An user
//
// swagger:model ClientDTO
type InvoiceDTO struct {
	// The id
	//
	// required: false
	Id uint `json:"id"`
	// The first name
	//
	// required: false
	Number      string    `json:"number"`
	Type        int       `json:"type"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	SubTotal    float32   `json:"subTotal"`
	VAT         float32   `json:"vat"`
	Total       float32   `json:"total"`
	ClientId    int64     `json:"clientId"`
}
