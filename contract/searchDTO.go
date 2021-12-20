package contract

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
	// The search tags
	//
	// required: false
	Tags []string `json:"tags"`
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
	// The search tags
	//
	// required: false
	Tags []string `json:"tags"`
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
	// The search tags
	//
	// required: false
	Tags []string `json:"tags"`
	// The participants
	//
	// required: false
	Participants []UserDTO
}