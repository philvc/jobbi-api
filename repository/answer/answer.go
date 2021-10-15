package answer_repository

import (
	"errors"

	contract "github.com/philvc/jobbi-api/contract"
	"github.com/philvc/jobbi-api/database/model"
	"gorm.io/gorm"
)

type AnswerRepository struct {
	database *gorm.DB
}

func Default(db *gorm.DB) AnswerRepository {
	return AnswerRepository{
		database: db,
	}
}

func (repository AnswerRepository) GetAnswersBySearchId(searchId uint) (*[]contract.AnswerDTO, error) {
	var answers []model.Answer
	var search model.Search

	if err := repository.database.Where("id = ?", searchId).First(&search).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	if err := repository.database.Model(&search).Association("Answers").Find(&answers); err != nil {
		return nil, err
	}

	AnswerDTOs := model.ToAnswerDTOs(answers)

	return &AnswerDTOs, nil
}

func (repository AnswerRepository) GetAnswerById(answerId string) (*contract.AnswerDTO, error) {
	var answer model.Answer

	if err := repository.database.Where("id = ?", answerId).First(&answer).Error; err != nil {
		return nil, errors.New("no Answer")
	}

	answerDTO := model.ToAnswerDTO(answer)

	return &answerDTO, nil
}

func (repository AnswerRepository) AddAnswer(answerDTO contract.AnswerDTO) (*contract.AnswerDTO, error) {

	answer := model.ToAnswer(answerDTO)

	if err := repository.database.Create(&answer).Error; err != nil {
		return nil, errors.New("failed to create Answer")
	}

	answerDTO = model.ToAnswerDTO(answer)

	return &answerDTO, nil
}

func (repository AnswerRepository) ModifyAnswer(answerDTO contract.AnswerDTO) (*contract.AnswerDTO, error) {

	answer := model.ToAnswer(answerDTO)

	repository.database.Model(&answer).Where("id = ?", answer.ID).Updates(map[string]interface{}{"type": answer.Type,
		"link": answer.Link, "description": answer.Description, "title": answer.Title})

	answerDTO = model.ToAnswerDTO(answer)

	return &answerDTO, nil
}
