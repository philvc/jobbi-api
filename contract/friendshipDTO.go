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
	// The userId
	//
	// required: false
	UserId string `json:"userId"`
	// The searchId
	//
	// required: false
	SearchId string `json:"searchId"`
	// The type
	//
	// required: false
	Type string `json:"type"`
}
