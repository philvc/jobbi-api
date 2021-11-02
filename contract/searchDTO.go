package contract

// An search
//
// swagger:model SearchDTO
type SearchDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// UserId
	//
	// required: false
	UserID string `json:"userId"`
}
