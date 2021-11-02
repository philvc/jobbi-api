package contract

// An offer
//
// swagger:model OfferDTO
type OfferDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: false
	Description string `json:"description"`
	// The title
	//
	// required: false
	Title string `json:"title"`
	// The link
	//
	// required: false
	Link string `json:"link"`
	// the search id
	//
	// required: false
	SearchID string `json:"searchId"`
	// the user id
	//
	// required: false
	UserID string `json:"userId"`
}
