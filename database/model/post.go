package model

import (
	"github.com/lib/pq"
	"github.com/philvc/jobbi-api/contract"
)

type Post struct {
	Base
	Description string
	Title       string
	UserID      string
	SearchID    string
	Type        string
	Comments    string
	PhoneNumber int64
	Email       string
	Tags        pq.StringArray `gorm:"type:text[]"`
	Url         string
	Company     string
	FirstName   string
	LastName    string
}

func ToPostDTO(post Post) contract.PostDTO {
	return contract.PostDTO{
		Id:          post.Base.ID,
		Description: post.Description,
		Title:       post.Title,
		UserID:      post.UserID,
		Tags:        post.Tags,
		Type:        post.Type,
		SearchID:    post.SearchID,
		FirstName:   post.FirstName,
		LastName:    post.LastName,
		Email:       post.Email,
		Company:     post.Company,
		Url:         post.Url,
		PhoneNumber: post.PhoneNumber,
	}
}

func ToPost(postDTO contract.PostDTO) Post {
	return Post{
		Base: Base{
			ID: postDTO.Id,
		},
		Description: postDTO.Description,
		Title:       postDTO.Title,
		UserID:      postDTO.UserID,
		Tags:        postDTO.Tags,
		Type:        postDTO.Type,
		SearchID:    postDTO.SearchID,
		FirstName:   postDTO.FirstName,
		LastName:    postDTO.LastName,
		Email:       postDTO.Email,
		Company:     postDTO.Company,
		Url:         postDTO.Url,
		PhoneNumber: postDTO.PhoneNumber,
	}
}

func ToPostDTOs(posts []Post) []contract.PostDTO {
	PostDtos := make([]contract.PostDTO, len(posts))

	for i, item := range posts {
		PostDtos[i] = ToPostDTO(item)
	}

	return PostDtos
}

func ToPosts(postsDTO []contract.PostDTO) []Post {
	Posts := make([]Post, len(postsDTO))

	for i, item := range postsDTO {
		Posts[i] = ToPost(item)
	}

	return Posts
}
