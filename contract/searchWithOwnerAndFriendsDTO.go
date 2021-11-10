package contract




// A Quest with owner and friends details
//
// swagger:model SearchWithOwnerAndFriendsDTO
type SearchWithOwnerAndFriendsDTO struct {
	// The search
	// 
	// required: false
	Search SearchDTO `json:"search"`
	// Friends
	//
	// required: false
	Friends []UserDTO `json:"friends"`
	// owner
	//
	// required: false
	Owner UserDTO `json:"owner"`
}