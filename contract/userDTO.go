package contract

// An user
//
// swagger:model UserDTO
type UserDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
	// The firstName
	//
	// required: false
	FirstName string `json:"firstName"`
	// The lastName
	//
	// required: false
	LastName string `json:"lastName"`
	// The email
	//
	// required: false
	Email string `json:"email"`
	// The externalId
	//
	// required: false
	ExternalId string `json:"externalId"`
}
