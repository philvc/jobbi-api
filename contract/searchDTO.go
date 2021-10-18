package contract

// An search
//
// swagger:model SearchDTO
type SearchDTO struct {
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
	// UserId
	//
	// required: false
	UserID uint `json:"userId"`
}