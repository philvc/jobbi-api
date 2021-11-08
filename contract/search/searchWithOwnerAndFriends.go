package contract

import "github.com/philvc/jobbi-api/contract"

type SearchWithOwnerAndFriends struct {
	// The search
	// 
	// required: true
	Search contract.SearchDTO `json:"search"`
	// Friends for search
	//
	// required: true
	Friends []contract.UserDTO `json:"friends"`
	// owner of search
	//
	// required: true
	Owner contract.UserDTO `json:"owner"`
}