package model

import "github.com/philvc/jobbi-api/contract"

type Comment struct {
	Base
	UserID   string
	SearchID string
	PostID   string
	Content  string
}

func ToComment(dto *contract.CommentDTO) Comment {
	return Comment{
		Base: Base{
			ID: dto.Id,
		},
		UserID:   dto.UserId,
		SearchID: dto.SearchId,
		Content:  dto.Content,
		PostID:   dto.PostId,
	}
}

func ToCommentDto(comment *Comment) *contract.CommentDTO {
	return &contract.CommentDTO{
		Id: comment.ID,
		UserId: comment.UserID,
		SearchId: comment.SearchID,
		Content: comment.Content,
	}
}