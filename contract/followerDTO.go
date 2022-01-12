package contract

// An friendship
//
// swagger:model FriendshipDTO
type FollowerDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The userId
	//
	// required: false
	UserId string `json:"userId"`
	// The searchId
	//
	// required: false
	SearchId string `json:"searchId"`
}