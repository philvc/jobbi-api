package contract

import "github.com/lib/pq"

// An search
//
// swagger:model SearchDTO
type SearchDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// UserId
	//
	// required: false
	UserID string `json:"userId"`
	// The search sector
	//
	// required: false
	Sector string `json:"sector"`
	// The search tags
	//
	// required: false
	Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
}

// An shared search
//
// swagger:model SharedSearchDTO
type SharedSearchDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// UserId
	//
	// required: false
	UserID string `json:"userId"`
	// The firstName
	//
	// required: false
	FirstName string `json:"firstName"`
	// The lastName
	//
	// required: false
	LastName string `json:"lastName"`
	// Avatar url
	//
	// required: false
	AvatarUrl string `json:"avatarUrl"`
	// The search sector
	//
	// required: false
	Sector string `json:"sector"`
	// The search tags
	//
	// required: false
	Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
}

// An followed search
//
// swagger:model FollowedSearchDTO
type FollowedSearchDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// UserId
	//
	// required: false
	UserID string `json:"userId"`
	// The firstName
	//
	// required: false
	FirstName string `json:"firstName"`
	// The lastName
	//
	// required: false
	LastName string `json:"lastName"`
	// Avatar url
	//
	// required: false
	AvatarUrl string `json:"avatarUrl"`
	// The search sector
	//
	// required: false
	Sector string `json:"sector"`
	// The search tags
	//
	// required: false
	Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
}

// My search
//
// swagger:model MySearchDTO
type MySearchDTO struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// The participants
	//
	// required: false
	Participants []UserDTO `json:"participants" gorm:"foreignKey:Id"` 
	// The search sector
	//
	// required: false
	Sector string `json:"sector"`
	// The search tags
	//
	// required: false
	Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
}
