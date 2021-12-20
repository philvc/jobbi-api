package contract

// A friend search
//
// swagger:model FriendSearchDTO
type FriendSearchDTO struct {
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
