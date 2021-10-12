package contract

// An answer
//
// swagger:model AnswerDTO
type AnswerDTO struct {
	// The id
	//
	// required: true
	Id uint `json:"id"`
	// The type
	//
	// required: true
	Type uint `json:"type"`
	// The link
	//
	// required: false
	Link          string `json:"link"`
	// The description
	//
	// required: false
	Description          string `json:"description"`
	// The title
	//
	// required: false
	Title          string `json:"title"`

}
