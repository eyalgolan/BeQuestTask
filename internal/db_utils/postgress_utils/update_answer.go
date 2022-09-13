package postgress_utils

import (
	"BeQuestPrep/internal/db_utils/postgress_utils/models"
	"BeQuestPrep/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (c *Client) updateHistoryUpdate(answer rest_utils.AnswerData) error {
	err := c.DB.Create(&models.Event{
		Event:   "update",
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
		return errors.Wrap(err, "add update event to history")
	}
	return nil
}
func (c *Client) UpdateAnswer(answer rest_utils.AnswerData) error {
	_, err := c.getLatestAnswer(answer.Key)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, "update answer")
	}
	err = c.DB.Create(&models.Answer{
		Key:     answer.Key,
		Value:   answer.Value,
		Deleted: false,
	}).Error
	if err != nil {
		return errors.Wrap(err, "update record")
	}
	return c.updateHistoryUpdate(answer)
}
