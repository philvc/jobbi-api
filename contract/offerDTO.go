package contract

// An offer
//
// swagger:model OfferDTO
type OfferDTO struct {
	// The id
	//
	// required: false
	Id uint `json:"id"`
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
	// required: true
	SearchID uint `json:"searchId"`
	// the user id
	//
	// required: false
	UserID uint `json:"userId"`
}
