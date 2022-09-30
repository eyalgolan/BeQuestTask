package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (c *DBClient) DeleteAnswer(answer rest_utils.AnswerData) error {
	latestAnswer, err := c.GetLatestAnswer(answer.Key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "Not found")
		}
		return errors.Wrap(err, "Unable to perform db action")
	}
	latestAnswer.Deleted = true
	err = c.DB.Save(&latestAnswer).Error
	if err != nil {
		return errors.Wrap(err, "delete answer")
	}
	return nil
}
