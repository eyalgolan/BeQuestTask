package postgress_utils

import (
	"BeQuestPrep/internal/db_utils"
	"BeQuestPrep/internal/db_utils/postgress_utils/models"
	"BeQuestPrep/internal/rest_utils"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (c *Client) updateHistoryCreate(answer rest_utils.AnswerData) error {
	err := c.DB.Create(&models.Event{
		Event:   "create",
		Key:     answer.Key,
		Value:   answer.Value,
		Deleted: false,
	}).Error
	if err != nil {
		err = c.DB.Where(
			"key = ? AND value = ? AND delete = ?",
			answer.Key,
			answer.Value,
			false).Delete(&models.Answer{}).Error
		if err != nil {
			return errors.Wrap(err, "rollback after unable to save to history")
		}
		return errors.Wrap(err, "add create event to history")
	}
	return nil
}

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
	return c.updateHistoryCreate(answer)
}
