package contract

// A comment
//
// swagger:model CommentForPostDto
type CommentForPostDto struct {
	// The id
	//
	// required: true
	Id string `json:"id"`
	// The content
	//
	// required: true
	Content string `json:"content"`
	// UserId
	//
	// required: true
	UserID string `json:"userId"`
	// The search id
	//
	// required: true
	SearchID string `json:"searchId"`
	// The post id
	//
	// required: true
	PostID string `json:"postId"`
	// The search id
	//
	// required: true
	FirstName string `json:"firstName"`
	// The post id
	//
	// required: true
	LastName string `json:"lastName"`
}

// A comment
//
// swagger:model CreateCommentResponseDto
type CreateCommentResponseDto struct {
	// The id
	//
	// required: true
	Id string `json:"id"`
	// The content
	//
	// required: true
	Content string `json:"content"`
	// UserId
	//
	// required: true
	UserID string `json:"userId"`
	// The search id
	//
	// required: true
	SearchID string `json:"searchId"`
	// The post id
	//
	// required: true
	PostID string `json:"postId"`
	// The search id
	//
	// required: true
	FirstName string `json:"firstName"`
	// The post id
	//
	// required: true
	LastName string `json:"lastName"`
}

// A comment
//
// swagger:model CommentUpdateRequestDto
type CommentUpdateRequestDto struct {
	// The content
	//
	// required: true
	Content string `json:"content"`
}

// A comment
//
// swagger:model CommentCreateRequestDto
type CommentCreateRequestDto struct {
	// The content
	//
	// required: true
	Content string `json:"content"`
}

// A comment
//
// swagger:model CommentUpdateDto
type CommentUpdateDto struct {
	// The id
	//
	// required: false
	Id string `json:"id"`
	// The content
	//
	// required: false
	Content string `json:"content"`
	// UserId
	//
	// required: true
	UserID string `json:"userId"`
	// The search id
	//
	// required: true
	SearchID string `json:"searchId"`
	// The post id
	//
	// required: true
	PostID string `json:"postId"`
	// The search id
	//
	// required: true

}

type CommentDTO struct {
	SearchId string
	UserId   string
	Content  string
	Id       string
	PostId   string
}
