package contract

// An company
//
// swagger:model CompanyDTO
type CompanyDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The link
	//
	// required: false
	Link string `json:"link"`
	// The description
	//
	// required: false
	Description string `json:"description"`
	// The title
	//
	// required: false
	Title string `json:"title"`
	// The UserID
	//
	// required: false
	UserID string `json:"userId"`
	// The SearchID
	//
	// required: false
	SearchID string `json:"searchId"`
}
