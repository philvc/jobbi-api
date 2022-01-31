package model

import "github.com/philvc/jobbi-api/contract"

type Comment struct {
	Base
	UserID   string
	SearchID string
	PostID   string
	Content  string
}

func ToComment(*contract.CommentDTO){
	
}