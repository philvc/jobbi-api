package contract

// An friendship
//
// swagger:model FriendshipDTO
type FriendshipDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
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
	UserId string `json:"userId"`
	// The searchId
	//
	// required: false
	SearchId string `json:"searchId"`

}
