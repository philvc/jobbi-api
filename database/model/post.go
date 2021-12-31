package model

import (
	"github.com/lib/pq"
	"github.com/philvc/jobbi-api/contract"
)

type Post struct {
	Base
	Description        string
	Title              string
	UserID             string
	SearchID           string
	Type               string
	Comments           string
	Tags               pq.StringArray `gorm:"type:text[]"`
	Url                string
	CompanyName        string
	CompanyEmail       string
	CompanyPhoneNumber int64
	CompanyAddress     string
	CompanyUrl         string
	ContactFirstName   string
	ContactLastName    string
	ContactPhoneNumber int64
	ContactEmail       string
}

func ToPostDTO(post Post) contract.PostDTO {
	return contract.PostDTO{
		Id:                 post.Base.ID,
		Description:        post.Description,
		Title:              post.Title,
		UserID:             post.UserID,
		Tags:               post.Tags,
		Type:               post.Type,
		SearchID:           post.SearchID,
		ContactFirstName:   post.ContactFirstName,
		ContactLastName:    post.ContactLastName,
		ContactEmail:       post.ContactEmail,
		CompanyName:        post.CompanyName,
		CompanyEmail:       post.CompanyEmail,
		CompanyPhoneNumber: post.CompanyPhoneNumber,
		CompanyAddress:     post.CompanyAddress,
		Url:                post.Url,
		ContactPhoneNumber: post.ContactPhoneNumber,
		CompanyUrl:         post.CompanyUrl,
	}
}

func ToPost(postDTO contract.PostDTO) Post {
	return Post{
		Base: Base{
			ID: postDTO.Id,
		},
		Description:        postDTO.Description,
		Title:              postDTO.Title,
		UserID:             postDTO.UserID,
		Tags:               postDTO.Tags,
		Type:               postDTO.Type,
		SearchID:           postDTO.SearchID,
		ContactFirstName:   postDTO.ContactFirstName,
		ContactLastName:    postDTO.ContactLastName,
		ContactEmail:       postDTO.ContactEmail,
		CompanyName:        postDTO.CompanyName,
		CompanyEmail:       postDTO.CompanyEmail,
		CompanyPhoneNumber: postDTO.CompanyPhoneNumber,
		CompanyAddress:     postDTO.CompanyAddress,
		CompanyUrl:         postDTO.CompanyUrl,
		Url:                postDTO.Url,
		ContactPhoneNumber: postDTO.ContactPhoneNumber,
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
