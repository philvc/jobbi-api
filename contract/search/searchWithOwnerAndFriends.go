package contract

import "github.com/philvc/jobbi-api/contract"



// A Quest with owner and friends details
//
// swagger:model SearchWithOwnerAndFriendsDTO
type SearchWithOwnerAndFriends struct {
	// The search
	// 
	// required: false
	Search contract.SearchDTO `json:"search"`
	// Friends
	//
	// required: false
	Friends []contract.UserDTO `json:"friends"`
	// owner
	//
	// required: false
	Owner contract.UserDTO `json:"owner"`
}