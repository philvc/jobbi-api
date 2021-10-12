package contract

// An user
//
// swagger:model OrganisationDTO
type OrganisationDTO struct {
	// The id
	//
	// required: false
	Id uint `json:"id"`
	// The name
	//
	// required: false
	Name string `json:"name"`
	// The slug
	//
	// required: false
	Slug string `json:"slug"`
	// The address
	//
	//required: false
	Address string `json:"address"`
	// The VAT Number
	//
	// required: false
	VATNumber string `json:"vatNumber"`
	// The owner Id
	//
	// required: false
	Owner uint `json:"userId"`

	Swift string `json:"swift"`

	BankAccount string `json:"bankAccount"`

	CompanyColor string `json:"companyColor"`

	ContactPhoneNumber string `json:"contactPhoneNumber"`

	ContactEmail string `json:"contactEmail"`
}
