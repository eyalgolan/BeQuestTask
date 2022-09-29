package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
)

func (c *Client) DeleteAnswer(answer rest_utils.AnswerData) error {
	latestAnswer, err := c.GetLatestAnswer(answer.Key)
	latestAnswer.Deleted = true
	err = c.DB.Save(&latestAnswer).Error
	if err != nil {
		return errors.Wrap(err, "delete answer")
	}
	return nil
}
