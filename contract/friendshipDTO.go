package contract

// An friendship
//
// swagger:model FriendshipDTO
type FriendshipDTO struct {
	// The id
	//
	// required: true
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
	// required: true
	LastName string `json:"lastName"`
	// The email
	//
	// required: true
	Email string `json:"email"`
	// The userId
	//
	// required: false
	UserId uint `json:"userId"`
	// The searchId
	//
	// required: true
	SearchId uint `json:"searchId"`

}
