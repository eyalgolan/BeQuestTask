package postgress_utils

import (
	"KeyValuePermStore/internal/db_utils"
	"KeyValuePermStore/internal/db_utils/postgress_utils/models"
	"KeyValuePermStore/internal/rest_utils"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (c *Client) CreateAnswer(answer rest_utils.AnswerData) error {
	var latestAnswer *models.Answer
	latestAnswer, err := c.getLatestAnswer(answer.Key)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "create answer")
	}
	if latestAnswer != nil {
		if latestAnswer.Value == answer.Value {
			return fmt.Errorf("input: %w", &db_utils.DuplicateAnswerErr{Answer: answer})
		}
	}

	err = c.DB.Create(&models.Answer{
		Key:     answer.Key,
		Value:   answer.Value,
		Deleted: false},
	).Error
	if err != nil {
		return errors.Wrap(err, "create record")
	}
	return c.CreateEvent(answer)
}
