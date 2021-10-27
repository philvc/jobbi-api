package contract

// An friendship
//
// swagger:model FriendshipDTO
type FriendshipDTO struct {
	// The id
	//
	// required: false
	Id uint `json:"id"`
	// The state
	//
	// required: false
	State uint `json:"state"`
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
	// The userId
	//
	// required: false
	UserId uint `json:"userId"`
	// The searchId
	//
	// required: false
	SearchId uint `json:"searchId"`

}
