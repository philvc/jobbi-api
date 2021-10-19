package contract

// An network
//
// swagger:model NetworkDTO
type NetworkDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
	// The description
	//
	// required: false
	Description string `json:"description"`
	// The firstName
	//
	// required: true
	FirstName string `json:"firstName"`
	// The lastName
	//
	// required: true
	LastName string `json:"lastName"`
	// The email
	//
	// required: false
	Email string `json:"email"`
	// The phoneNumber
	//
	// required: false
	PhoneNumber string `json:"phoneNumber"`
	// The link
	//
	// required: false
	Link string `json:"link"`
	// The search id
	//
	// required: true
	SearchID uint `json:"searchId"`
	// The user id
	//
	// required: false
	UserID uint `json:"userId"`
}
