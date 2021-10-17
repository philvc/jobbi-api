package contract

// An offer
//
// swagger:model OfferDTO
type OfferDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The title
	//
	// required: true
	Title string `json:"title"`
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
