package postgress_utils

import (
	"github.com/pkg/errors"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
)

func (c *Client) CreateAnswer(answer rest_utils.AnswerData) error {
	//var latestAnswer *models.Answer
	//latestAnswer, err := c.GetLatestAnswer(answer.Key)
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return errors.Wrap(err, "create answer")
	//}
	//if latestAnswer != nil {
	//	if latestAnswer.Value == answer.Value {
	//		return fmt.Errorf("input: %w", &db_utils.DuplicateAnswerErr{Answer: answer})
	//	}
	//}

	err := c.DB.Create(&models.Answer{
		Key:     answer.Key,
		Value:   answer.Value,
		Deleted: false},
	).Error
	if err != nil {
		return errors.Wrap(err, "create record")
	}
	return nil
}
