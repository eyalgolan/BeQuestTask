package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
)

func (c *Client) DeleteEvent(answer rest_utils.AnswerData) error {
	err := c.DB.Create(&models.Event{
		Event:   "delete",
		Key:     answer.Key,
		Value:   answer.Value,
		Deleted: true,
	}).Error
	if err != nil {
		var answerToRollback models.Answer
		err = c.DB.Last(
			&answerToRollback,
			"key = ? AND value = ? and deleted = ?", answer.Key, answer.Value, true).Error
		if err != nil {
			return errors.Wrap(err, "rollback after unable to save to history")
		}
		answerToRollback.Deleted = false
		err = c.DB.Save(&answerToRollback).Error
		if err != nil {
			return errors.Wrap(err, "rollback after unable to save to history")
		}
		return errors.Wrap(err, "add delete event to history")
	}
	return nil
}
