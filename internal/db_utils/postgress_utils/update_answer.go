package postgress_utils

import (
	"BeQuestPrep/internal/db_utils/postgress_utils/models"
	"BeQuestPrep/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

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
