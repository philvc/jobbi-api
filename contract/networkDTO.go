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
	// required: true
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
	// required: true
	Email string `json:"email"`
	// The phoneNumber
	//
	// required: true
	PhoneNumber string `json:"phoneNumber"`
	// The link
	//
	// required: true
	Link string `json:"link"`
	// the search id
	//
	// required: true
	SearchID uint `json:"searchId"`
	// the user id
	//
	// required: true
	UserID uint `json:"userId"`
}
