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
	// The contact firstName
	//
	// required: false
	ContactFirstName string `json:"contactFirstName"`
	// The contact lastName
	//
	// required: false
	ContactLastName string `json:"contactLastName"`
	// The contact email
	//
	// required: false
	ContactEmail string `json:"contactEmail"`
	// The company name
	//
	// required: false
	CompanyName string `json:"companyName"`
	// The company email
	//
	// required: false
	CompanyEmail string `json:"companyEmail"`
	// The company phone number
	//
	// required: false
	CompanyPhoneNumber int64 `json:"companyPhoneNumber"`
	// The company address
	//
	// required: false
	CompanyAddress string `json:"companyAddress"`
	// The company url
	//
	// required: false
	CompanyUrl string `json:"companyUrl"`
	// url
	//
	// required: false
	Url string `json:"url"`
	// The contact phoneNumber
	//
	// required: false
	ContactPhoneNumber int64 `json:"contactPhoneNumber"`
}

// An post
//
// swagger:model PostDTOBySearchId
type PostDTOBySearchId struct {
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
	// The contact firstName
	//
	// required: false
	ContactFirstName string `json:"contactFirstName"`
	// The contact lastName
	//
	// required: false
	ContactLastName string `json:"contactLastName"`
	// The contact email
	//
	// required: false
	ContactEmail string `json:"contactEmail"`
	// The user firstName
	//
	// required: false
	UserFirstName string `json:"userFirstName"`
	// The user  lastName
	//
	// required: false
	UserLastName string `json:"userLastName"`
	// The user email
	//
	// required: false
	UserEmail string `json:"userEmail"`
	// The company name
	//
	// required: false
	CompanyName string `json:"companyName"`
	// The company email
	//
	// required: false
	CompanyEmail string `json:"companyEmail"`
	// The company phone number
	//
	// required: false
	CompanyPhoneNumber int64 `json:"companyPhoneNumber"`
	// The company address
	//
	// required: false
	CompanyAddress string `json:"companyAddress"`
	// The company url
	//
	// required: false
	CompanyUrl string `json:"companyUrl"`
	// url
	//
	// required: false
	Url string `json:"url"`
	// The contact phoneNumber
	//
	// required: false
	ContactPhoneNumber int64 `json:"contactPhoneNumber"`
}

// An post
//
// swagger:model AddPostRequestDTO
type AddPostRequestDTO struct {
	// The title
	//
	// required: true
	Title string `json:"title"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The search type
	//
	// required: true
	Type string `json:"type"`
	// url
	//
	// required: false
	Url string `json:"url"`
}

// An post
//
// swagger:model AddPostResponseDTO
type AddPostResponseDTO struct {
	// The id
	//
	// required: true
	Id string `json:"id"`
	// The title
	//
	// required: true
	Title string `json:"title"`
	// The description
	//
	// required: true
	Description string `json:"description"`
	// The search type
	//
	// required: true
	Type string `json:"type"`
	// url
	//
	// required: false
	Url string `json:"url"`
	// The search id
	//
	// required: true
	SearchID string `json:"searchId"`
	// UserId
	//
	// required: true
	UserID string `json:"userId"`
	// The user firstName
	//
	// required: false
	UserFirstName string `json:"userFirstName"`
	// The user  lastName
	//
	// required: false
	UserLastName string `json:"userLastName"`
	// The user email
	//
	// required: false
	UserEmail string `json:"userEmail"`
}
