package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (c *DBClient) UpdateAnswer(answer rest_utils.AnswerData) error {
	_, err := c.GetLatestAnswer(answer.Key)
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
	return nil
}
