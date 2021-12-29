package contract

import "github.com/lib/pq"

// An post
//
// swagger:model PostDTO
type PostDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: false
	Description string `json:"description"`
	// The title
	//
	// required: false
	Title string `json:"title"`
	// UserId
	//
	// required: true
	UserID string `json:"userId"`
	// The search sector
	//
	// required: true
	SearchID string `json:"searchId"`
	// The search type
	//
	// required: true
	Type string `json:"type"`
	// The search tags
	//
	// required: false
	Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
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
	// The company
	//
	// required: false
	Company string `json:"company"`
	// url
	//
	// required: false
	Url string `json:"url"`
	// The phoneNumber
	//
	// required: false
	PhoneNumber int64 `json:"phoneNumber"`
}