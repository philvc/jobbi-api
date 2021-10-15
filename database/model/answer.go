package model

import (
	"github.com/philvc/jobbi-api/contract"
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Type     uint
	Description      string
	Title         string
	Link    string
	SearchID uint
	FriendshipID uint
}

func ToAnswerDTO(answer Answer) contract.AnswerDTO {
	return contract.AnswerDTO{
		Id:         answer.ID,
		Type:  answer.Type,
		Link:   answer.Link,
		Description:      answer.Description,
		Title: answer.Title,
	}
}

func ToAnswer(answerDTO contract.AnswerDTO) Answer {
	return Answer{
		Model: gorm.Model{
			ID: answerDTO.Id,
		},
		Type:  answerDTO.Type,
		Link:   answerDTO.Link,
		Description:      answerDTO.Description,
		Title: answerDTO.Title,
	}
}

func ToAnswerDTOs(answers []Answer) []contract.AnswerDTO {
	AnswerDTOs := make([]contract.AnswerDTO, len(answers))

	for i, item := range answers {
		AnswerDTOs[i] = ToAnswerDTO(item)
	}

	return AnswerDTOs
}

func ToAnswers(answerDTOs []contract.AnswerDTO) []Answer {
	Answers := make([]Answer, len(answerDTOs))

	for i, item := range answerDTOs {
		Answers[i] = ToAnswer(item)
	}

	return Answers
}
