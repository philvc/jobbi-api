package contract

// An company
//
// swagger:model CompanyDTO
type CompanyDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
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
	UserID uint `json:"userId"`
	// The SearchID
	//
	// required: false
	SearchID uint `json:"searchId"`
}
