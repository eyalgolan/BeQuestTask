package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
)

func (c *DBClient) CreateAnswer(answer rest_utils.AnswerData) error {
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
