package contract

// A follower
//
// swagger:model FollowerDTO
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
